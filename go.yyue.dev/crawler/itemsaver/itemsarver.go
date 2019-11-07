package itemsaver

import (
	"context"
	"encoding/gob"
	"fmt"

	"github.com/micro/go-micro"
	"go.yyue.dev/common/types"
	"go.yyue.dev/common/utils"
	"go.yyue.dev/crawler/engine"
	"go.yyue.dev/crawler/proto"
	databaseproto "go.yyue.dev/datamanage/proto"
)

func New() chan engine.Item {
	item := make(chan engine.Item)
	gob.Register(types.HouseInfo{})

	go func() {
		count := 0
		srv := micro.NewService(
			micro.Name("database.client"),
		)
		srv.Init()

		database := databaseproto.NewDatabaseServiceClient("database", srv.Client())
		rep, err := database.Connect(context.TODO(), &databaseproto.ConnectRequest{Database: "development"})
		utils.PanicErr(err)
		fmt.Println(rep)

		for {
			result := <-item
			count++
			go func() {
				house := result.(proto.House)
				message, err := database.InsertHouse(context.TODO(), &databaseproto.InsertHouseRequest{Database: "development", House: &house})

				if err != nil {
					fmt.Println("error:", err, message)
				} else {
					fmt.Println("finally: ", message)
				}
			}()
		}
	}()
	return item
}
