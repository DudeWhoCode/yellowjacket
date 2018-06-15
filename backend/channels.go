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

func GetResponseChan() (ResponseChan chan RawResponse) {
	ResponseChan = make(chan RawResponse)
	return ResponseChan
}
