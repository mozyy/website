package main

import (
	"fmt"
	"net/rpc"

	"yyue.dev/common/message"
	"yyue.dev/common/utils"
)

type User struct {
	Name string `db:"name"`
}

func main() {
	datamanageURL := utils.GetConfig().DatamanageURL
	client, err := rpc.DialHTTP("tcp", datamanageURL)
	utils.PanicErr(err)
	defer client.Close()

	message := &message.Message{}
	// user := &[]User{
	// 	{"name"},
	// }
	user := "test"
	err = client.Call("Query.Show", &user, message)
	utils.PanicErr(err)
	fmt.Println(message)
}
