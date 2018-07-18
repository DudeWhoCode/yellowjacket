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
		for i, m := range wasp.Tasks {
			fmt.Println("Counter: ", i)
			start := time.Now()
			fmt.Println("Executing one of the task: ", wasp.Tasks)
			fmt.Println("Executing the task: ", m)
			response := wasp.ExecuteOne(m)
			elapsed := time.Since(start)
			fmt.Println("Elapsed time: ", elapsed)
			responseStruct := struct {
				ResponseHeader *http.Response
				ResponseTime   string
			}{
				response,
				elapsed.String(),
			}
			// fmt.Println("Response Struct Created ---- : \n", responseStruct)
			s.OutChan <- responseStruct
			fmt.Println("After pushing to channel :::::::::::::::::::::;")
		}
		time.Sleep(time.Second)
	}
}

// Collect listens to a channel to gather ther responses from StartAttack
func (s *swarm) Collect() {
	fmt.Println("Listening for logs")
	for resp := range s.OutChan {
		if resp.ResponseHeader.StatusCode == 200 {
			Agg.SumReq++
		} else {
			Agg.SumFails++
		}
		b.Messages <- fmt.Sprintf("%d, %d, %s", Agg.SumReq, Agg.SumFails, resp.ResponseTime)
		fmt.Println(resp.ResponseHeader.StatusCode, " took ", resp.ResponseTime)
	}
}
