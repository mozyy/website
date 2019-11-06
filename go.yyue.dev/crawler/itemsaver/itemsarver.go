package itemsaver

import (
	"encoding/gob"
	"fmt"

	"go.yyue.dev/common/types"
	"go.yyue.dev/crawler/engine"
	"go.yyue.dev/crawler/parser/lianjia"
)

func New() chan engine.Item {
	item := make(chan engine.Item)
	gob.Register(types.HouseInfo{})

	go func() {
		count := 0
		// datamanageURL := utils.GetConfig().DatamanageURL
		// client, err := rpc.DialHTTP("tcp", datamanageURL)
		// utils.PanicErr(err)
		// defer client.Close()
		// msg := &message.Message{}
		// err = client.Call("Query.ConnectDatabase", "development", msg)
		// utils.PanicErr(err)
		// fmt.Println("ConnectDatabase: ", msg)

		for {
			result := <-item
			count++
			if houseInfo, ok := result.(lianjia.HouseInfo); ok {
				fmt.Printf("got result: %v, count: %d\n", houseInfo, count)

			}
			// go func() {
			// 	message := &message.Message{}
			// 	fmt.Println("start: ", message)
			// 	dbo := types.DBOperater{"development", "house_info", result}
			// 	err := client.Call("Query.Insert", dbo, message)
			// 	if err != nil {
			// 		fmt.Println("error:", err, message)
			// 	}
			// 	fmt.Println("finally: ", message)
			// }()
		}
	}()
	return item
}
