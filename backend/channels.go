package backend

import "net/http"

type Swarm struct {
	Inputs  SwarmInput
	OutChan chan RawResponse
}

type SwarmInput struct {
	HatchRate int `json:"hatch_rate"`
	Wasps     int `json:"wasps"`
}

type AggrResponse struct {
	SumReq   int
	SumFails int
}

type RawResponse struct {
	ResponseHeader *http.Response
	ResponseTime   string
}

var ResponseChan RawResponse

var s *Swarm

func GetSwarm(input SwarmInput) (s *Swarm) {
	if s == nil {
		s = &Swarm{
			input,
			make(chan RawResponse),
		}
	}
	return
}
func GetResponseChan() (ResponseChan chan RawResponse) {
	ResponseChan = make(chan RawResponse)
	return ResponseChan
}
