package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/DudeWhoCode/yellowjacket/backend"
)

type Data struct {
	Host string
}

// Ping returns the availability of the webserver
func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

// StartSwarm is used to kick off the load test
func StartSwarm(w http.ResponseWriter, r *http.Request) {
	// Decode the post payload
	var a backend.SwarmInput
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&a); err != nil {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Unable to parse json"))
		return
	}
	defer r.Body.Close()
	swarm := backend.GetSwarm()
	swarm.WebInputs(a.Wasps, a.HatchRate)
	swarm.SetChan(make(chan backend.RawResponse))
	go swarm.Collect()
	go swarm.CreateSwarm()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

// Home serves the home page of the app
func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Read in the template with SSE JavaScript code.
	t, err := template.ParseFiles("templates/index.html")
	if err != nil {
		log.Fatal("Error parsing your template.")

	}
	data := Data{
		backend.TargetHost,
	}
	// Render the template, writing to `w`.
	t.Execute(w, data)
}

func StopSwarm(w http.ResponseWriter, r *http.Request) {
	swarm := backend.GetSwarm()
	swarm.StopFlag = true
	swarm.NumReq = 0
	swarm.NumFail = 0
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}
