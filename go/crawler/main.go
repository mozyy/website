package main

import (
	"fmt"

	"yyue.dev/crawler/engine"
	parser "yyue.dev/crawler/parser/lianjia"
	"yyue.dev/crawler/scheduler"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()
	item := make(chan engine.Item)
	entry := engine.Engine{
		Scheduler:   scheduler.New(),
		WorkerCount: 15,
		ItemSaver:   item,
	}
	go func() {
		count := 0

		for {
			result := <-item
			count++
			fmt.Printf("got result: %s, count: %d\n", result, count)
		}
	}()
	request := engine.Request{
		URL: "https://cd.lianjia.com/ershoufang/",
		// URL: "https://cd.lianjia.com/ershoufang/106103087618.html",
		Parser: parser.City,
	}
	entry.Run(request)
}
