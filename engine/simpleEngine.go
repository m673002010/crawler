package engine

import (
	"crawler/model"
	"crawler/worker"
	"log"
)

type SimpleEngine struct{}

func (e SimpleEngine) Run(seeds ...model.Request) {
	var requests []model.Request

	// 请求加入请求队列
	requests = append(requests, seeds...)

	// 循环执行队列任务
	for len(requests) > 0 {
		log.Println("req length: ", len(requests))
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker.Worker(r)

		// 忽略错误继续爬取下一个url
		if err != nil {
			log.Printf("Fetcher error url: %s err: %v", r.Url, err)
			continue
		}

		requests = append(requests, parserResult.Requests...)
	}

	// 队列清空
	log.Println("end")
}
