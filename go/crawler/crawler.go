package main

import "crawler/engin"

func main() {
	entry := engin.Request{
		URL:    "https://cd.lianjia.com/ershoufang/",
		Parser: engin.NilParser,
	}
	entry.Run()
}
