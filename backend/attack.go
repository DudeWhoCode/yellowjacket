package backend

import (
	"fmt"
	"net/http"
	"time"
)

var Agg AggrResponse

// CreateSwarm takes in the waspsCount(users count) and hatchRate(spwan rate) to
// create a swarm of concurrent users to attack. hatchRate tells how many concurrent
// users it has to spawn every second.
func (s *Swarm) CreateSwarm() {
	hatchRate := s.Inputs.HatchRate
	waspsCount := s.Inputs.Wasps
	fmt.Printf("Starting %d wasps at the rate of %d/sec \n", waspsCount, hatchRate)
	remainingWasps := waspsCount
	for i := 0; i < waspsCount; i++ {
		for j := 0; j < hatchRate; j++ {
			go s.StartAttack("https://battlelog.battlefield.com/bf4/servers/")
			remainingWasps--
		}
		if remainingWasps == 0 {
			fmt.Printf("Created %d wasps \n", waspsCount)
			break
		}
		time.Sleep(time.Second)
	}
}

// StartAttack will hit the given url indefinetly until it's stopped by the user
func (s *Swarm) StartAttack(url string) {

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
		s.OutChan <- responseStruct
		fmt.Println("Sent to channel")
	}
}

// Collect listens to a channel to gather ther responses from StartAttack
func (s *Swarm) Collect() {
	fmt.Println("Listening for logs")
	for resp := range s.OutChan {
		if resp.ResponseHeader.StatusCode == 200 {
			Agg.SumReq++
		} else {
			Agg.SumFails++
		}
		b.Messages <- fmt.Sprintf("%d, %d, %s", Agg.SumReq, Agg.SumFails, resp.ResponseTime)
		// fmt.Println(resp.ResponseHeader.StatusCode, " took ", resp.ResponseTime)
	}
}
