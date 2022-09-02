package main

import (
	"fmt"
	"net/http"
)

type response struct {
	url    string
	status string
}

func main() {
	results := make(map[string]string)
	channel := make(chan response)
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
	}

	for _, url := range urls {
		go hitURL(url, channel)
	}

	for i := 0; i < len(urls); i++ {
		response := <-channel
		results[response.url] = response.status
	}

	for url, status := range results {
		fmt.Println(url, status)
	}
}

func hitURL(url string, channel chan<- response) {
	fmt.Println("Checking:", url)

	resp, err := http.Get(url)
	var status string
	if err != nil || resp.StatusCode >= 400 {
		status = "OK"
	} else {
		status = "FAILED"
	}
	channel <- response{url: url, status: status}
}
