package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/DudeWhoCode/yellowjacket/backend"
)

func Ping(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}

func AttackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside attack handler")
	var a backend.Attack
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
	pipe := backend.InitializeCollect()
	go backend.CollectLogs(pipe)
	go backend.CreateSwarm(swarmRate, wasps, pipe)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
}

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
