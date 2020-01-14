package itemsaver

import (
	"database/sql"
	"encoding/gob"
	"log"

	"go.yyue.dev/common/message"
	"go.yyue.dev/common/types"
	"go.yyue.dev/crawler/engine"
	"go.yyue.dev/crawler/proto"
	"go.yyue.dev/database"
)

func New() chan engine.Item {
	item := make(chan engine.Item)
	gob.Register(types.HouseInfo{})

	go func() {
		db, err := database.GetDB("crawler")
		if err != nil {
			return
		}
		for {
			result := <-item
			go sarverHandler(db, result)
		}
	}()
	return item
}

var (
	count    = 0
	total    = 0
	errCount = 0
)

func sarverHandler(db *sql.DB, result engine.Item) {
	switch value := result.(type) {
	case proto.House:
		err := InsertHouse(db, value)

		if err != nil {
			log.Printf("saver error: %s, request: %v\n", err, value.GetHouseInfo().GetUrl())
			// reSaver(result, database, err, message, value.GetHouseInfo().GetUrl())
		}
	case proto.HouseSummary:
		err := InsertHouseSummary(db, value)

		if err != nil {
			log.Printf("saver error: %s, request: %v\n", err, value.GetUrl())
			// reSaver(result, database, err, message, value.GetUrl())
		}
	case int:
		total += value
		log.Printf("page: %d, count: %d, total: %d\n", count, value, total)
	}
}

func reSaver(result engine.Item, err error, message *message.Message, detail string) {
	// if errCount > 200 {
	// 	panic(fmt.Sprintf("errorHouse: %s, message: %s, url: %s\n", err, message, detail))
	// }
	// <-time.NewTimer(time.Second).C
	// errCount++
	log.Printf("saver error: %s, request: %v, resave...\n", err, detail)
	// sarverHandler(result, database)
}
