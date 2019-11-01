package main

import (
	"github.com/PuerkitoBio/goquery"
	"yyue.dev/crawler/engine"
	"yyue.dev/crawler/itemsaver"
	parser "yyue.dev/crawler/parser/lianjia"
	"yyue.dev/crawler/scheduler"
)

func main() {
	entry := engine.Engine{
		Scheduler:   scheduler.New(),
		WorkerCount: 15,
		ItemSaver:   itemsaver.New(),
	}
	request := engine.Request{
		URL: "https://cd.lianjia.com/ershoufang/",
		Parser: parser.City,
		// URL: "https://cd.lianjia.com/ershoufang/106103655131.html",
		// Parser: func(q *goquery.Document) engine.Result {
		// 	return parser.House(q, "https://cd.lianjia.com/ershoufang/106103655131.html")
		// },
	}
	entry.Run(request)
}
