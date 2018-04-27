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
	go collectLogs()
	startServer("8000")
	fmt.Println("Listening on port 8000: ")
}

func AttackHandler(w http.ResponseWriter, r *http.Request) {
	//decoder := json.NewDecoder(r.Body)
	//fmt.Println("Got post form : ", r.Body)
	//var a attack
	//if err := decoder.Decode(&a); err != nil {
	//	w.WriteHeader(http.StatusNotAcceptable)
	//	w.Write([]byte("Unable to parse json"))
	//	return
	//}

	r.ParseForm()
	swarmRate := 5 //strconv.Atoi(r.FormValue("swarm"))
	wasps, _ := strconv.Atoi(r.FormValue("wasps"))
	defer r.Body.Close()
	fmt.Println("wasp count: ", wasps, swarmRate)
	go createSwarm(swarmRate, wasps)
}



func createSwarm(swarmRate int, wasps int) {
	fmt.Printf("Starting %d wasps at the rate of %d/sec \n", wasps, swarmRate)
	remainingWasps := wasps
	for i := 0; i < wasps; i++ {
		fmt.Println("Wasping")
		for j:=0; j<swarmRate; j++ {
			println("Into j")
			go Attack("http://www.google.com")
			remainingWasps--
		}
		if remainingWasps == 0 {
			fmt.Printf("Created %d wasps \n", wasps)
			break
		}
		fmt.Println("Before time.sleep")
		time.Sleep(time.Second)
	}
}

func Attack(url string) {

	for ;; {
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

func collectLogs() {
	for resp := range responses {
		continue
		fmt.Println(resp.ResponseHeader.StatusCode, " took ", resp.ResponseTime)
	}
	fmt.Println("Blocking receiver")
}

func startServer(port string) {
	http.HandleFunc("/attack", AttackHandler) // set router
	port = ":" + port
	err := http.ListenAndServe(port, nil)  // set listen port
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
}