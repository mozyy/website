package main

import (
	"fmt"
	"net/rpc"

	"github.com/mozyy/website/go/message"
	"github.com/mozyy/website/go/user/model"

	"github.com/mozyy/utils"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:5000")
	utils.PanicErr(err)
	defer client.Close()

	message := &message.Message{}
	user := model.User{10, "rpc client", 18381335182, "123"}
	err = client.Call("User.Regist", user, message)
	utils.PanicErr(err)
	fmt.Println(message)
}
