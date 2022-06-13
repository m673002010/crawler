package main

import (
	"crawler/engine"
	"crawler/myParser"
)

func main() {
	// 爬虫开始run
	engine.Run(engine.Request{
		Url:        "https://movie.douban.com/top250", // url
		ParserFunc: myParser.ParseTop250,              // url对应的解析函数
	})

	// test_unit.Test_parseMovie()
}
