package engine

import "go.yyue.dev/crawler/fetcher"

type Scheduler interface {
	Submit(Request)
	GetWorkerChan() chan Request
	WorkerReady(chan Request)
}

// Run is the entry point for engin
func (engine *Engine) Run(seed ...Request) {
	out := make(chan Result)
	for i := 0; i < engine.WorkerCount; i++ {
		createWorker(engine.Scheduler.GetWorkerChan(), out, engine.Scheduler)
	}
	for _, request := range seed {
		engine.Scheduler.Submit(request)
	}
	for {
		result := <-out
		// save items of result
		for _, item := range result.Items {
			go func(item Item) { engine.ItemSaver <- item }(item)
		}
		// next requests
		for _, request := range result.Requests {
			if checkDuplicate(request.URL) {
				engine.Scheduler.Submit(request)
			}
		}
	}
}
func createWorker(in chan Request, out chan Result, scheduler Scheduler) {
	go func() {
		for {
			scheduler.WorkerReady(in)
			request := <-in
			result := worker(request)
			out <- result
		}
	}()
}
func worker(request Request) Result {
	b := fetcher.Fetch(request.URL)
	result := request.Parser(b)
	return result
}

var visitsMaps = make(map[string]bool)

// 验证url是否重复
func checkDuplicate(url string) bool {
	// 不存在会返回 false
	if visitsMaps[url] {
		return false
	}
	visitsMaps[url] = true
	return true
}
