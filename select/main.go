package main

import (
	"net/http"
)

// 你被要求编写一个叫做 WebsiteRacer 的函数，用来对比请求两个 URL 来「比赛」，并返回先响应的 URL。如果两个 URL 在 10 秒内都未返回结果，那么应该返回一个 error。

func Racer(a, b string) (winner string) {
	select {
	case <-ping(a):
		return a
	case <-ping(b):
		return b
	}
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}

func main() {

}

/*
myvar := <- ch 这是一个阻塞调用，myvar 在等待值发送给channel
select 则允许同时在多个channel 等待，第一个发送值的channel 胜出，case 重的代码会被执行
*/
