package itemsaver

import (
	"encoding/gob"
	"fmt"
	"net/rpc"

	"yyue.dev/common/message"
	"yyue.dev/common/utils"
	"yyue.dev/crawler/engine"
	parser "yyue.dev/crawler/parser/lianjia"
)

func New() chan engine.Item {
	item := make(chan engine.Item)
	gob.Register(parser.HouseInfo{})

	go func() {
		count := 0
		datamanageURL := utils.GetConfig().DatamanageURL
		client, err := rpc.DialHTTP("tcp", datamanageURL)
		utils.PanicErr(err)
		defer client.Close()
		for {
			result := <-item
			count++
			fmt.Printf("got result: %s, count: %d\n", result, count)
			go func() {
				message := &message.Message{}

				err := client.Call("Query.Insert", &result, message)
				utils.PanicErr(err) // TODO: handler error
				fmt.Println(message)
			}()
		}
	}()
	return item
}
