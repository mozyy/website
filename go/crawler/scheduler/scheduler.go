package scheduler

import "yyue.dev/crawler/engine"

type Scheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

// New is create scheduler and init scheduler channel listener.
func New() *Scheduler {
	scheduler := &Scheduler{
		requestChan: make(chan engine.Request),
		workerChan:  make(chan chan engine.Request),
	}

	go func() {
		requestQueue := []engine.Request{}
		workerQueue := []chan engine.Request{}
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQueue) > 0 && len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			select {
			case request := <-scheduler.requestChan:
				requestQueue = append(requestQueue, request)
			case worker := <-scheduler.workerChan:
				workerQueue = append(workerQueue, worker)
			case activeWorker <- activeRequest:
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()

	return scheduler
}

func (s *Scheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

func (s *Scheduler) WorkerReady() {

}
