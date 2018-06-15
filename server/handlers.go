package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/DudeWhoCode/yellowjacket/backend"
)

// Ping returns the availability of the webserver
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

// StartAttack is used to kick off the load test
func StartAttack(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside attack handler")
	// Decode the post payload
	var a backend.Attack
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Unable to parse json"))
		return
	}
	defer r.Body.Close()
	fmt.Println(a)
	swarmRate := a.Swarms
	wasps := a.Wasps
	fmt.Println("wasp count: ", wasps, swarmRate)
	pipe := backend.InitializeCollect()
	go backend.CollectLogs(pipe)
	go backend.CreateSwarm(swarmRate, wasps, pipe)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// MainHandler serves the home page of the app
func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Read in the template with our SSE JavaScript code.
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("WTF dude, error parsing your template.")

	}

	// Render the template, writing to `w`.
	t.Execute(w, "YellowJacket")
}
