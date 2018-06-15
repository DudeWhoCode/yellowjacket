package backend

import "net/http"

type Swarm struct {
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

func GetResponseChan() (ResponseChan chan RawResponse) {
	ResponseChan = make(chan RawResponse)
	return ResponseChan
}
