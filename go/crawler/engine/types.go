package engine

type Engine struct {
	Scheduler Scheduler
	ItemSaver chan Item
}

// Request is engin request
type Request struct {
	URL    string
	Parser func([]byte) Result
}

// Result is engin result
type Result struct {
	Requests []Request
	Items    []Item
}

// Items is engin result
type Item struct {
	URL string

	Perloads interface{}
}

// ParserResult is engin result
// type ParserResult struct {
// 	Requests []Request
// 	Items    []interface{}
// }

// NilParser is engin nil Parser, for not complete Parser
func NilParser(b []byte) Result {
	return Result{}
}
