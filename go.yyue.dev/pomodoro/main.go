package main

import (
	"fmt"
	"net/rpc"

	"go.yyue.dev/common/message"
	"go.yyue.dev/common/utils"
)

// User is test
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
