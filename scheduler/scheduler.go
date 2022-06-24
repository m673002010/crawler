package scheduler

import (
	"crawler/model"
)

type Scheduler interface {
	Distribute(model.Request)
}

type SimpleScheduler struct{}

func (s *SimpleScheduler) Distribute(r model.Request, in chan model.Request) {
	in <- r
}
