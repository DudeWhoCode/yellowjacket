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

// StartSwarm is used to kick off the load test
func StartSwarm(w http.ResponseWriter, r *http.Request) {
	// Decode the post payload
	var a backend.Swarm
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Unable to parse json"))
		return
	}
	defer r.Body.Close()
	hatchRate := a.HatchRate
	wasps := a.Wasps
	pipe := backend.GetResponseChan()
	go backend.CollectLogs(pipe)
	go backend.CreateSwarm(hatchRate, wasps, pipe)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// MainHandler serves the home page of the app
func MainHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Read in the template with SSE JavaScript code.
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("WTF dude, error parsing your template.")

	}

	// Render the template, writing to `w`.
	t.Execute(w, "YellowJacket")
}
