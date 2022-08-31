package main

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

var (
	errorRequestFailed = errors.New("Request failed")
)

func main() {
	// urls := []string{
	// 	"https://www.airbnb.com/",
	// 	"https://www.google.com/",
	// 	"https://www.amazon.com/",
	// 	"https://www.reddit.com/",
	// 	"https://www.google.com/",
	// 	"https://soundcloud.com/",
	// 	"https://www.facebook.com/",
	// 	"https://www.instagram.com/",
	// }
	// var results = make(map[string]string)

	// for _, url := range urls {
	// 	result := "OK"
	// 	err := hitURL(url)

	// 	if err != nil {
	// 		result = "FAILED"
	// 	}
	// 	results[url] = result
	// 	fmt.Println(url, result)
	// }

	//* initiating, type of info that will be channeled
	channel := make(chan string)
	people := [2]string{"Tony", "Red"}
	for _, person := range people {
		go isPresent(person, channel)
	}
	// channel의 메시지를 받을 때까지 main함수는 기다린다
	// Promise.all[]과 비슷하네
	//! goroutine의 결과를 받는 receiving channel은 블로킹 싱크 작업이다
	// 받을때까지 멈춥니다
	//! 블로킹 속성을 해결하기 위한 채널 버퍼링
	for i := 0; i < len(people); i++ {
		fmt.Print("Blocking, hence stops here until receiving ")
		fmt.Println(<-channel)
	}
	// go count("Tony")
	// go count("Red")
	//* main은 Goroutine을 기다리지 않는다
	//* main이 꺼지지 않도록 count("Bella")는 살아있어야 한다
	//* 매번 time.sleep()을 걸어줄 수 없다.
	// 이렇게 하지 않고 goroutine이 끝날 때까지 main이 기다릴 수 있는, 다시 말해 goroutine과 main이 소통할 수 있는
	// Channel
	// Goroutine ↔ Goroutine ↔ Main
	// count("Bella")

	fmt.Println("Main function is finished")
}

func hitURL(url string) error {
	fmt.Println("Checking:", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		return errorRequestFailed
	}
	defer resp.Body.Close()
	return nil
}

func count(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is present", i)
		time.Sleep(time.Second)
	}
}

func isPresent(person string, channel chan string) {
	time.Sleep(time.Second * 2)
	// channel에다 결과값을 주는 방법
	channel <- person + " is present"
}
