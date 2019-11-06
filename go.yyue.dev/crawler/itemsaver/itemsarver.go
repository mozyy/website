package itemsaver

import (
	"encoding/gob"
	"fmt"
	"net/rpc"

	"go.yyue.dev/common/message"
	"go.yyue.dev/common/types"
	"go.yyue.dev/common/utils"
	"go.yyue.dev/crawler/engine"
)

func New() chan engine.Item {
	item := make(chan engine.Item)
	gob.Register(types.HouseInfo{})

	go func() {
		count := 0
		datamanageURL := utils.GetConfig().DatamanageURL
		client, err := rpc.DialHTTP("tcp", datamanageURL)
		utils.PanicErr(err)
		defer client.Close()
		msg := &message.Message{}
		err = client.Call("Query.ConnectDatabase", "development", msg)
		utils.PanicErr(err)
		fmt.Println("ConnectDatabase: ", msg)

		for {
			result := <-item
			count++
			fmt.Printf("got result: %s, count: %d\n", result, count)
			go func() {
				message := &message.Message{}
				fmt.Println("start: ", message)
				dbo := types.DBOperater{"development", "house_info", result}
				err := client.Call("Query.Insert", dbo, message)
				if err != nil {
					fmt.Println("error:", err, message)
				}
				fmt.Println("finally: ", message)
			}()
		}
	}()
	return item
}
