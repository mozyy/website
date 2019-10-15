package engine

import "github.com/PuerkitoBio/goquery"

type Engine struct {
	Scheduler   Scheduler
	ItemSaver   chan Item
	WorkerCount int
}

// Request is engin request
type Request struct {
	URL    string
	Parser func(*goquery.Document) Result
}

// Result is engin result
type Result struct {
	Requests []Request
	Items    []Item
}

// Items is engin result
type Item interface{}

// ParserResult is engin result
// type ParserResult struct {
// 	Requests []Request
// 	Items    []interface{}
// }

// NilParser is engin nil Parser, for not complete Parser
func NilParser(q *goquery.Document) Result {
	return Result{}
}
