package itemsaver

import (
	"context"
	"encoding/gob"
	"log"

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
		// all := map[string]string{}
		database := databaseproto.NewDatabaseServiceClient("database", srv.Client())
		_, err := database.Connect(context.TODO(), &databaseproto.ConnectRequest{Database: "development"})
		utils.PanicErr(err)

		for {
			result := <-item
			count++
			go func() {
				house := result.(proto.House)
				// info := house.GetHouseInfo()
				// url := info.GetUrl()
				// no := info.GetHouseNo()
				// title := info.GetTitle()
				// log.Printf("count: %d, url: %s, house: %s\n", count, url, title)
				// if _, ok := all[no]; ok {
				// 	log.Printf("mutile no: %s", no)
				// } else {
				// 	all[no] = title
				// }
				message, err := database.InsertHouse(context.TODO(), &databaseproto.InsertHouseRequest{Database: "development", House: &house})

				if err != nil {
					log.Printf("error: %s, message: %s, url: %s\n", err, message, house.GetHouseInfo().GetUrl())
				}
			}()
		}
	}()
	return item
}
