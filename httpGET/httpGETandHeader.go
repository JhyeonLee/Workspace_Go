package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	// Request 객체 생성
	req, err := http.NewRequest("GET", "http://csharp.tips/feed/rss", nil)
	if err != nil {
		panic(err)
	}

	// 필요한 헤더 추가 가능
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:69.0) Gecko/20100101 firefox/69.0")

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// 결과 출력
	bytes, _ := ioutil.ReadAll(resp.Body)
	str := string(bytes) // 바이트를 문자열로
	fmt.Println(str)
}
