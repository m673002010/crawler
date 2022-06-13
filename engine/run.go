package engine

import (
	"crawler/fetcher"
	"log"
)

func Run(seeds ...Request) {
	var requests []Request

	// 请求加入请求队列
	requests = append(requests, seeds...)

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		log.Println("req length: ", len(requests))

		// 抓取页面数据
		body, err := fetcher.Fetch(r.Url)

		// 忽略错误继续爬取下一个url
		if err != nil {
			log.Printf("Fetcher error url: %s err: %v", r.Url, err)
			continue
		}

		// 解析页面数据，新请求入队
		parserResult := r.ParserFunc(body)
		requests = append(requests, parserResult.Requests...)
	}

	log.Println("end")
}
