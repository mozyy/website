package websocket

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"yyue.dev/common/utils"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type message struct {
	UID    uid         `json:"uid"`
	Target uid         `json:"target"`
	Kind   string      `json:"kind"`
	Value  interface{} `json:"value"`
}
type userInfoStruct struct {
	uid     uid
	channel string
}

type uid = int

type user struct {
	uid  uid
	conn *websocket.Conn
}

var currentUID uid = 10000
var channels = map[string][]user{}

// Handler websocket handler for rtc
func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	var userInfo userInfoStruct

	closeHandler := func(code int, text string) error {
		log.Printf("close message: %d %s \n", code, text)
		for i, uid := range channels[userInfo.channel] {
			if uid.uid == userInfo.uid {
				uids := channels[userInfo.channel]
				uids = append(uids[:i], uids[i+1:]...)
				channels[userInfo.channel] = uids
				log.Printf("current users: %v\n", uids)
				break
			}
		}
		return &websocket.CloseError{Code: code, Text: text}
	}
	conn.SetCloseHandler(closeHandler)

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Printf("received message: %d %s \n", messageType, string(p))
		var msg message
		err = json.Unmarshal(p, &msg)
		if err != nil {
			log.Println(err)
			return
		}

		switch msg.Kind {
		case "join":
			if userInfo.uid != 0 {
				sendMessage(conn, message{Kind: "join-failure", Value: "uid already exists"})
				return
			}
			channel := msg.Value.(map[string]interface{})["channel"].(string)
			users := channels[channel]
			userInfo = userInfoStruct{
				uid:     currentUID,
				channel: channel,
			}
			currentUID++
			users = append(users, user{uid: userInfo.uid, conn: conn})
			channels[channel] = users
			sendMessage(conn, message{
				Kind:   "join-success",
				Target: userInfo.uid,
			})
			sendMessage(conn, message{
				Kind:   "userlist",
				UID:    userInfo.uid,
				Target: userInfo.uid,
				Value:  users,
			})
			sendMessageOther(channel, userInfo.uid, message{
				Kind: "user-joined",
				UID:  userInfo.uid,
			})
		case "create-offer":
			fallthrough
		case "create-answer":
			fallthrough
		case "new-ice-candidate":
			fallthrough
		case "hang-up":
			sendMessageByTarget(userInfo.channel, msg)

		default:
			log.Println("received unknown message: ", msg.Kind)
			sendMessageByTarget(userInfo.channel, msg)

			// if err := conn.WriteMessage(messageType, p); err != nil {
			// 	log.Println(err)
			// 	return
			// }
		}
	}
}

func sendMessageByTarget(channel string, msg message) {
	target := msg.Target
	if target != 0 {
		sendMessageByUID(channel, target, msg)
	} else {
		sendMessageByChannel(channel, msg)
	}
}

func sendMessageByChannel(channel string, msg message) {
	for _, user := range channels[channel] {
		sendMessage(user.conn, msg)
	}
}

func sendMessageByUID(channel string, uid uid, msg message) {
	conn := getConnByUID(channel, uid)
	sendMessage(conn, msg)
}

func sendMessageOther(channel string, uid uid, msg message) {
	for _, user := range channels[channel] {
		if uid != user.uid {
			sendMessage(user.conn, msg)
		}
	}
}

func getConnByUID(channel string, uid uid) *websocket.Conn {
	for _, user := range channels[channel] {
		if user.uid == uid {
			return user.conn
		}
	}
	utils.PanicErr(errors.New("not found conn"))
	return nil
}

func sendMessage(conn *websocket.Conn, msg message) {
	reply, err := json.Marshal(msg)
	utils.LogErr(err)
	err = conn.WriteMessage(websocket.BinaryMessage, reply)
	utils.LogErr(err)
}
