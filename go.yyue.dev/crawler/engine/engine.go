package engine

import (
	"errors"
	"fmt"
	"log"
	"time"

	"go.yyue.dev/crawler/fetcher"
)

// Scheduler is crawler engine scheduler
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
			// if checkDuplicate(request.URL) {
			engine.Scheduler.Submit(request)
			// }
		}
	}
}

func createWorker(in chan Request, out chan Result, scheduler Scheduler) {
	var fetchErrorCount = 0
	go func() {
		for {
			scheduler.WorkerReady(in)
			request := <-in
		worker:
			result, err := worker(request)
			if err != nil {
				if errors.Is(err, ErrListLarge) {
					log.Printf("worker error: %s, request: %v\n", err, request)
				} else {
					if fetchErrorCount > 50 {
						panic("fetch error too mach, please check")
					}
					<-time.NewTimer(time.Second).C
					fetchErrorCount++
					log.Printf("worker error: %s, request: %v, rework...\n", err, request)
					goto worker
				}

			}
			out <- result

		}
	}()
}

func worker(request Request) (Result, error) {
	b, err := fetcher.Fetch(request.URL)
	if err != nil {
		return Result{}, fmt.Errorf("fetch error: %s", err)
	}
	return request.Parser(b)
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
