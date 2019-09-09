package websocket

import (
	"encoding/json"
	"log"
	"net/http"

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
	kind  string
	value interface{}
}

type user = int

var channels = map[string][]user{}

func Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		msg := message{}
		json.Unmarshal(p, &msg)

		switch msg.kind {
		case "join":
			channel := msg.value.(map[string]string)["channel"]
			users := channels[channel]
			users = append(users, len(users))
			channels[channel] = users
			reply, err := json.Marshal(message{kind: "userList", value: users})
			if err != nil {
				log.Println(err)
				return
			}
			if err := conn.WriteMessage(messageType, reply); err != nil {
				log.Println(err)
				return
			}
		default:
			if err := conn.WriteMessage(messageType, p); err != nil {
				log.Println(err)
				return
			}
		}

	}

}
