package main

import (
	"github.com/PuerkitoBio/goquery"
	"go.yyue.dev/crawler/engine"
	"go.yyue.dev/crawler/itemsaver"
	"go.yyue.dev/crawler/parser/lianjia"
	"go.yyue.dev/crawler/scheduler"
)

func main() {
	entry := engine.Engine{
		Scheduler:   scheduler.New(),
		WorkerCount: 15,
		ItemSaver:   itemsaver.New(),
	}
	request := engine.Request{
		// URL:    "https://cd.lianjia.com/ershoufang/",
		// Parser: lianjia.City,
		URL: "https://cd.lianjia.com/ershoufang/106103655131.html",
		Parser: func(q *goquery.Document) engine.Result {
			return lianjia.House(q, "https://cd.lianjia.com/ershoufang/106103655131.html")
		},
	}
	entry.Run(request)
}
