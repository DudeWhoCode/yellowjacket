package main

import (
	"log"
	"net/http"

	"github.com/DudeWhoCode/yellowjacket/backend"
	"github.com/DudeWhoCode/yellowjacket/server"
	"github.com/rs/cors"
)

func main() {
	rt := server.NewRouter()
	handler := cors.Default().Handler(rt)
	// Setup the file server to serve the static site with css and js
	fs := http.FileServer(http.Dir("templates/"))
	rt.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", fs))
	// Start the SSE broker
	b := backend.GetBroker()
	b.Start()
	log.Fatal(http.ListenAndServe(":8001", handler))
}
