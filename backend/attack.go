package backend

import (
	"fmt"
	"net/http"
	"time"
)

var responses chan struct {
	ResponseHeader *http.Response
	ResponseTime   string
}

func CreateSwarm(swarmRate int, wasps int) {
	fmt.Printf("Starting %d wasps at the rate of %d/sec \n", wasps, swarmRate)
	remainingWasps := wasps
	for i := 0; i < wasps; i++ {
		fmt.Println("Wasping")
		for j := 0; j < swarmRate; j++ {
			println("Spawning...")
			go Attack("http://www.google.com")
			remainingWasps--
		}
		if remainingWasps == 0 {
			fmt.Printf("Created %d wasps \n", wasps)
			break
		}
		time.Sleep(time.Second)
	}
}

func Attack(url string) {

	for {
		start := time.Now()
		response, _ := http.Get(url)
		elapsed := time.Since(start)

		responseStruct := struct {
			ResponseHeader *http.Response
			ResponseTime   string
		}{
			response,
			elapsed.String(),
		}
		responses <- responseStruct
	}
}

func collectLogs() {
	for resp := range responses {
		fmt.Println(resp.ResponseHeader.StatusCode, " took ", resp.ResponseTime)
	}
	fmt.Println("Blocking receiver")
}
