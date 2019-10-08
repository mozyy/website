package main

import (
	"yyue.dev/crawler/engine"
	"yyue.dev/crawler/scheduler"
)

func main() {
	entry := engine.Engine{
		Scheduler: scheduler.New(),
	}
	// entry := engine.Engine{
	// 	URL:    "https://cd.lianjia.com/ershoufang/",
	// 	Parser: engine.NilParser,
	// }
	entry.Run()
}
