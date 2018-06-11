package backend

import "net/http"

type AggrResponse struct {
	SumReq   int
	SumFails int
}
type Attack struct {
	Swarms int `json:"swarms"`
	Wasps  int `json:"wasps"`
}

type RawResponse struct {
	ResponseHeader *http.Response
	ResponseTime   string
}

var ResponseChan RawResponse
var AggrResponseChan RawResponse

func InitializeCollect() (ResponseChan chan RawResponse) {
	ResponseChan = make(chan RawResponse)
	return ResponseChan
}

func InitializeAggregate() (AggrResponseChan chan RawResponse) {
	AggrResponseChan = make(chan RawResponse)
	return AggrResponseChan
}
