package engine

import (
	"crawler/model"
	"crawler/saver"
	"crawler/scheduler"
	"crawler/worker"
)

type ConcurrentEngine struct{}

func (e *ConcurrentEngine) Run(seeds ...model.Request) {
	sch := scheduler.SimpleScheduler{}

	in := make(chan model.Request)
	out := make(chan model.ParserResult)

	for i := 0; i < 3; i++ {
		worker.CreateWorker(in, out)
	}

	for _, r := range seeds {
		go sch.Distribute(r, in)
	}

	count := 0
	for {
		result := <-out
		count++
		// log.Println("res:", count, result)
		go saver.Save(result)

		for _, req := range result.Requests {
			go sch.Distribute(req, in)
		}
	}
}
