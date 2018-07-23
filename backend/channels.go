package backend

import "net/http"

type SwarmInput struct {
	HatchRate int `json:"hatch_rate"`
	Wasps     int `json:"wasps"`
}

type AggrResponse struct {
	SumReq     int
	SumFails   int
	SumLatency float64
}

type RawResponse struct {
	ResponseHeader *http.Response
	ResponseTime   float64
}

var ResponseChan RawResponse

func GetResponseChan() (ResponseChan chan RawResponse) {
	ResponseChan = make(chan RawResponse)
	return ResponseChan
}
