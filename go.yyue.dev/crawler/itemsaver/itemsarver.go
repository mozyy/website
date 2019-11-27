package itemsaver

import (
	"context"
	"encoding/gob"
	"log"
	"time"

	"github.com/micro/go-micro"
	"go.yyue.dev/common/message"
	"go.yyue.dev/common/types"
	"go.yyue.dev/crawler/engine"
	"go.yyue.dev/crawler/proto"
	databaseproto "go.yyue.dev/datamanage/proto"
)

func New() chan engine.Item {
	item := make(chan engine.Item)
	gob.Register(types.HouseInfo{})

	go func() {

		srv := micro.NewService(
			micro.Name("database.client"),
		)
		srv.Init()
		database := databaseproto.NewDatabaseServiceClient("database", srv.Client())
		database.Connect(context.TODO(), &databaseproto.ConnectRequest{Database: "development"})
		// utils.PanicErr(err)

		for {
			result := <-item
			go sarverHandler(result, &database)
		}
	}()
	return item
}

var (
	count    = 0
	total    = 0
	errCount = 0
)

func sarverHandler(result engine.Item, database *databaseproto.DatabaseServiceClient) {
	switch value := result.(type) {
	case proto.House:
		message, err := (*database).InsertHouse(context.TODO(),
			&databaseproto.InsertHouseRequest{Database: "development", House: &value})

		if err != nil {
			reSaver(result, database, err, message, value.GetHouseInfo().GetUrl())
		}
	case proto.HouseSummary:
		message, err := (*database).InsertHouseSummary(context.TODO(),
			&databaseproto.InsertHouseSummaryRequest{Database: "development", House: &value})

		if err != nil {
			reSaver(result, database, err, message, value.GetUrl())
		}
	case int:
		total += value
		log.Printf("page: %d, count: %d, total: %d\n", count, value, total)
	}
}

func reSaver(result engine.Item, database *databaseproto.DatabaseServiceClient, err error, message *message.Message, detail string) {
	if errCount > 200 {
		log.Printf("errorHouse: %s, message: %s, url: %s\n", err, message, detail)
		return
	}
	<-time.NewTimer(time.Second).C
	errCount++
	log.Printf("saver error: %s, request: %v, resave...\n", err, detail)
	sarverHandler(result, database)
}
