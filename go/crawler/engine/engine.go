package engine

type Scheduler interface {
	Submit(Request)
	WorkerReady()
}

// Run is the entry point for engin
func (engine *Engine) Run(seed ...Request) {
	out := make(chan Result)
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

func worker() {
	// b := fetcher.Fetch(request.URL)

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
