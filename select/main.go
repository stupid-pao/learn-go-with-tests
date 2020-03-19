package main

import (
	"net/http"
	"time"
)

// 你被要求编写一个叫做 WebsiteRacer 的函数，用来对比请求两个 URL 来「比赛」，并返回先响应的 URL。如果两个 URL 在 10 秒内都未返回结果，那么应该返回一个 error。

func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}
	return b
}

func main() {

}
