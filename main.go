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
	fs := http.FileServer(http.Dir("templates/"))
	rt.PathPrefix("/templates/").Handler(http.StripPrefix("/templates/", fs))
	// http.Handle("/", rt)
	b := backend.GetBroker()
	// Start processing events
	b.Start()

	// Make b the HTTP handler for "/events/".  It can do
	// this because it has a ServeHTTP method.  That method
	// is called in a separate goroutine for each
	// request to "/events/".
	rt.Handle("/events/", b)

	log.Fatal(http.ListenAndServe(":8001", handler))
}
