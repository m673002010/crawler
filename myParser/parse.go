package myParser

import (
	"crawler/model"
	"log"
	"regexp"
	"strconv"
)

func ParseTop250(contents []byte) model.ParserResult {
	htmlStr := string(contents)

	// 获取电影id
	idReg := regexp.MustCompile(`<a href="https://movie.douban.com/subject/(.*?)/">`)
	movieIds := idReg.FindAllStringSubmatch(htmlStr, -1)

	// 获取电影名
	titleReg := regexp.MustCompile(`<img width="100" alt="(.*?)" src`)
	titles := titleReg.FindAllStringSubmatch(htmlStr, -1)

	result := model.ParserResult{}

	// 电影名、电影详情url、解析函数ParserFunc
	for _, title := range titles {
		result.Items = append(result.Items, title[1])
	}

	for _, id := range movieIds {
		reqUrl := "https://movie.douban.com/subject/" + id[1]
		result.Requests = append(result.Requests, model.Request{
			Url:        reqUrl,
			ParserFunc: ParseMovie,
		})
	}

	return result
}

func ParseMovie(contents []byte) model.ParserResult {
	htmlStr := string(contents)

	movie := model.Movie{}

	// 标题
	titleReg := regexp.MustCompile(`<meta property="og:title" content="(.*?)" />`)
	movie.Title = titleReg.FindAllStringSubmatch(htmlStr, -1)[0][1]

	// 时长
	durationReg := regexp.MustCompile(`<meta property="video:duration" content="(.*?)" />`)
	duration, err := strconv.Atoi(durationReg.FindAllStringSubmatch(htmlStr, -1)[0][1])
	if err != nil {
		log.Printf("duration error: %v", err)
	}
	movie.Duration = duration

	// 评分
	ratingReg := regexp.MustCompile(`<strong class="ll rating_num" property="v:average">(.*?)</strong>`)
	rating, err := strconv.ParseFloat(ratingReg.FindAllStringSubmatch(htmlStr, -1)[0][1], 64)
	if err != nil {
		log.Printf("rating error: %v", err)
	}
	movie.Rating = rating

	// 演员表
	actorReg := regexp.MustCompile(`<meta property="video:actor" content="(.*?)" />`)
	actors := actorReg.FindAllStringSubmatch(htmlStr, -1)
	for _, a := range actors {
		movie.Actors = append(movie.Actors, a[1])
	}

	// 导演
	directorReg := regexp.MustCompile(`<meta property="video:director" content="(.*?)" />`)
	directors := directorReg.FindAllStringSubmatch(htmlStr, -1)
	for _, d := range directors {
		movie.Director = append(movie.Director, d[1])
	}

	// 链接
	urlReg := regexp.MustCompile(`<meta property="og:url" content="(.*?)" />`)
	movie.Url = urlReg.FindAllStringSubmatch(htmlStr, -1)[0][1]

	// 图片
	imageReg := regexp.MustCompile(`<meta property="og:image" content="(.*?)" />`)
	movie.Image = imageReg.FindAllStringSubmatch(htmlStr, -1)[0][1]

	// 返回电影详情信息
	result := model.ParserResult{
		Items: []interface{}{movie},
	}

	return result
}
