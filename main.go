package main

import (
	"fmt"
	"net/http"
	"time"
	"strconv"
)

var responses chan struct {
	ResponseHeader *http.Response
	ResponseTime   string
}

type attack struct {
	Swarms int	`json:"swarms"`
	Wasps int	`json:"wasps"`
}

func main() {
	fmt.Println("yellowjacket welcomes you. buzzzzz..")
	responses = make(chan struct {
		ResponseHeader *http.Response
		ResponseTime   string
	})
	go func() {
		for resp := range responses {
			fmt.Println(resp.ResponseHeader.StatusCode, " took ", resp.ResponseTime)
		}
		fmt.Println("Blocking receiver")
	}()
	http.HandleFunc("/attack", AttackHander) // set router
	err := http.ListenAndServe(":8000", nil) // set listen port
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
	fmt.Println("Listening on port 8000: ")
}

func AttackHander(w http.ResponseWriter, r *http.Request) {
	//decoder := json.NewDecoder(r.Body)
	//fmt.Println("Got post form : ", r.Body)
	//var a attack
	//if err := decoder.Decode(&a); err != nil {
	//	w.WriteHeader(http.StatusNotAcceptable)
	//	w.Write([]byte("Unable to parse json"))
	//	return
	//}

	r.ParseForm()
	swarms, _ := strconv.Atoi(r.FormValue("swarms"))
	wasps, _ := strconv.Atoi(r.FormValue("wasps"))
	defer r.Body.Close()
	fmt.Println("wasp count: ", wasps, swarms)
	go createSwarm(swarms, wasps)
}

func createSwarm(swarms int, wasps int) {
	fmt.Println("Into CreateSwarm")
	for i := 0; i < wasps; i++ {
		fmt.Println("Wasping")
		AttackGet(swarms, "http://www.google.com")
	}
}

func AttackGet(swarm int, url string) {
	for i := 0; i < swarm; i++ {
		fmt.Println("Attacking : ", swarm, url)
		start := time.Now()
		response, _ := http.Get(url)
		elapsed := time.Since(start)

		responseStruct := struct {
			ResponseHeader *http.Response
			ResponseTime string

		}{
			response,
			elapsed.String(),
		}
		responses <- responseStruct
	}
}
