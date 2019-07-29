package main

import (
	"fmt"
	"net/rpc"

	"github.com/mozyy/website/go/message"
	"github.com/mozyy/website/go/user/model"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:5000")
	utils.PanicErr(err)
	defer client.Close()

	message := &message.Message{}
	user := model.User{2, "rpc client", 18381335182, "123"}
	err = client.Call("User.Login", user, message)
	fmt.Println(message)
	utils.PanicErr(err)
	fmt.Println(message)
}
