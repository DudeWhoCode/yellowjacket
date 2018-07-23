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
func (s *swarm) CreateSwarm() {
	hatchRate := s.HatchRate
	waspsCount := s.WaspCount
	wasp := s.Wasp
	wasp.GetMethods()
	fmt.Println("Tasks in hand: ", wasp.Tasks)
	fmt.Printf("Starting %d wasps at the rate of %d/sec \n", waspsCount, hatchRate)
	remainingWasps := waspsCount
	for i := 0; i < waspsCount; i++ {
		for j := 0; j < hatchRate; j++ {
			go s.StartAttack(wasp)
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
func (s *swarm) StartAttack(wasp *UserBehaviour) {

	for {
		if s.StopFlag == true {
			break
		}
		for i, m := range wasp.Tasks {
			if s.StopFlag == true {
				break
			}
			fmt.Println("Counter: ", i)
			start := time.Now()
			fmt.Println("Executing one of the task: ", wasp.Tasks)
			fmt.Println("Executing the task: ", m)
			response := wasp.ExecuteOne(m)
			elapsed := time.Since(start)
			fmt.Println("Elapsed time: ", elapsed)
			responseStruct := struct {
				ResponseHeader *http.Response
				ResponseTime   float64
			}{
				response,
				elapsed.Seconds(),
			}
			s.OutChan <- responseStruct
		}
		time.Sleep(time.Second)
	}
}

// Collect listens to a channel to gather ther responses from StartAttack
func (s *swarm) Collect() {
	fmt.Println("Listening for logs")
	for resp := range s.OutChan {
		if s.StopFlag == true {
			Agg.SumReq = 0
			Agg.SumFails = 0
			b.Messages <- fmt.Sprintf("%d, %d, %s", Agg.SumReq, Agg.SumFails, "0.0")
			continue
		}
		Agg.SumLatency += resp.ResponseTime
		if resp.ResponseHeader.StatusCode == 200 {
			Agg.SumReq++
		} else {
			Agg.SumFails++
		}
		avgLatency := Agg.SumLatency / float64(Agg.SumReq+Agg.SumFails)
		b.Messages <- fmt.Sprintf("%d, %d, %f", Agg.SumReq, Agg.SumFails, avgLatency)
		fmt.Println(resp.ResponseHeader.StatusCode, " took ", resp.ResponseTime)
	}
}
