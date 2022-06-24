package main

import (
	"crawler/engine"
	"crawler/model"
	"crawler/myParser"
)

func main() {
	// 单机爬虫开始run
	// engine.SimpleEngine{}.Run(model.Request{
	// 	Url:        "https://movie.douban.com/top250", // url
	// 	ParserFunc: myParser.ParseTop250,              // url对应的解析函数
	// })

	// cfg := elasticsearch.Config{
	// 	Addresses: []string{
	// 		"http://192.168.43.144:9200",
	// 	},
	// }

	// es, err := elasticsearch.NewClient(cfg)
	// if err != nil {
	// 	log.Fatalf("Error getting response: %s", err)
	// }
	// res, _ := es.Info()

	// log.Println("es info:", res)

	// 并发爬虫
	seeds := []model.Request{}
	// for i := 0; i < 10; i++ {
	// 	seeds = append(seeds, model.Request{
	// 		Url:        "https://movie.douban.com/top250?start=" + strconv.Itoa(i*25), // url
	// 		ParserFunc: myParser.ParseTop250,                                          // url对应的解析函数
	// 	})
	// }
	seeds = append(seeds, model.Request{
		Url:        "https://movie.douban.com/top250?start=0", // url
		ParserFunc: myParser.ParseTop250,                      // url对应的解析函数
	})

	e := &engine.ConcurrentEngine{}
	e.Run(seeds...)

	// test_unit.Test_parseMovie()
	// test_unit.Test_func()
}
