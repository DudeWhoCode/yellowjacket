package backend

import (
	"fmt"
	"net/http"
	"time"
)

var Agg AggrResponse

func CreateSwarm(swarmRate int, wasps int, pipe chan RawResponse) {
	fmt.Printf("Starting %d wasps at the rate of %d/sec \n", wasps, swarmRate)
	remainingWasps := wasps
	for i := 0; i < wasps; i++ {
		// fmt.Println("Wasping")
		for j := 0; j < swarmRate; j++ {
			// println("Spawning...")
			go StartAttack("https://battlelog.battlefield.com/bf4/servers/", pipe)
			remainingWasps--
		}
		if remainingWasps == 0 {
			fmt.Printf("Created %d wasps \n", wasps)
			break
		}
		time.Sleep(time.Second)
	}
}

func StartAttack(url string, pipe chan RawResponse) {

	for {
		start := time.Now()
		response, _ := http.Get(url)
		elapsed := time.Since(start)
		time.Sleep(time.Second)
		responseStruct := struct {
			ResponseHeader *http.Response
			ResponseTime   string
		}{
			response,
			elapsed.String(),
		}
		fmt.Println("Before sending to channel")
		go CollectLogs(pipe)
		pipe <- responseStruct
		fmt.Println("Sent to channel")
	}
}

func CollectLogs(pipe chan RawResponse) {
	fmt.Println("Listening for logs")
	for resp := range pipe {
		fmt.Println("STATUS CODE: :::: :: ", resp.ResponseHeader.StatusCode)
		if resp.ResponseHeader.StatusCode == 200 {
			Agg.SumReq++
		} else {
			Agg.SumFails++
		}
		b.Messages <- fmt.Sprintf("%d, %d, %s", Agg.SumReq, Agg.SumFails, resp.ResponseTime)
		fmt.Println("Sent to SSE")
		// fmt.Println(resp.ResponseHeader.StatusCode, " took ", resp.ResponseTime)
	}
	fmt.Println("Exiting the collectLogs ::::::::::::")
}
