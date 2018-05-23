package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/DudeWhoCode/yellowjacket/backend"
)

type attack struct {
	Swarms int `json:"swarms"`
	Wasps  int `json:"wasps"`
}

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func AttackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside attack handler")
	var a attack
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Unable to parse json"))
		return
	}
	defer r.Body.Close()
	fmt.Println(a)
	swarmRate := a.Swarms //strconv.Atoi(r.FormValue("toggle_box"))
	wasps := a.Wasps
	fmt.Println("wasp count: ", wasps, swarmRate)
	pipe := backend.Initialize()
	go backend.CollectLogs(pipe)
	go backend.CreateSwarm(swarmRate, wasps, pipe)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
