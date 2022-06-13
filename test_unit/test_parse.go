package test_unit

import (
	"crawler/myParser"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Test_parseMovie() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://movie.douban.com/subject/1292052/", nil)
	if err != nil {
		fmt.Println("http err", err)
	}

	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/79.0.3945.130 Safari/537.36 OPR/66.0.3515.115")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("client err", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("wrong statusCode", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("read err", err)
	}

	myParser.ParseMovie(body)
}
