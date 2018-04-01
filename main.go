package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

var responses chan struct {
	ResponseHeader *http.Response
	ResponseTime   string
}

func main() {
	wasps, _ := strconv.Atoi(os.Args[1])
	swarms, _ := strconv.Atoi(os.Args[2])

	fmt.Println("yellowjacket welcomes you. buzzzzz..")
	responses = make(chan struct {
		ResponseHeader *http.Response
		ResponseTime   string
	})
	for i := 0; i < wasps; i++ {
		go Get(swarms, "http://www.google.com")
	}

	func() {
		for resp := range responses {
			fmt.Println(resp.ResponseHeader.StatusCode, " took ", resp.ResponseTime)
		}
		fmt.Println("Blocking receiver")
	}()
}

func Get(swarm int, url string) {
	for i := 0; i < swarm; i++ {
		start := time.Now()
		response, _ := http.Get(url)
		elapsed := time.Since(start)

		responseStruct := struct {
			ResponseHeader *http.Response
			ResponseTime string

		}{
			response,
			elapsed.String(),
		}
		responses <- responseStruct
	}
}
