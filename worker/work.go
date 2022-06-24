package worker

import (
	"crawler/fetcher"
	"crawler/model"
)

func Worker(r model.Request) (model.ParserResult, error) {
	// 抓取页面数据
	body, err := fetcher.Fetch(r.Url)

	if err != nil {
		return model.ParserResult{}, err
	}

	// 解析页面数据
	return r.ParserFunc(body), nil
}

func CreateWorker(in chan model.Request, out chan model.ParserResult) {
	go func() {
		for {
			r := <-in
			result, err := Worker(r)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
