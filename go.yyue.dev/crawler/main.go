package main

import (
	"os"

	"go.yyue.dev/common/utils"
	"go.yyue.dev/crawler/engine"
	"go.yyue.dev/crawler/itemsaver"
	"go.yyue.dev/crawler/parser/lianjia"
	"go.yyue.dev/crawler/scheduler"
)

func main() {
	if len(os.Args) > 1 {
		utils.WaitForIt(os.Args[1])
	}
	entry := engine.Engine{
		Scheduler:   scheduler.New(),
		WorkerCount: 15,
		ItemSaver:   itemsaver.New(),
	}
	request := engine.Request{
		URL:    "https://cd.lianjia.com/ershoufang/",
		Parser: lianjia.City,
		// URL: "https://d.lianjia.com/ershoufang/106102678685.html",
		// Parser: func(q *goquery.Document) engine.Result {
		// 	return lianjia.House(q, "https://cd.lianjia.com/ershoufang/106102678685.html")
		// },
	}
	entry.Run(request)
}
