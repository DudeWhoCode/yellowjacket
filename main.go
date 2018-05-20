package main

import (
	"log"
	"net/http"

	"github.com/DudeWhoCode/yellowjacket/server"
	"github.com/rs/cors"
)

func main() {
	rt := server.NewRouter()
	handler := cors.Default().Handler(rt)
	log.Fatal(http.ListenAndServe(":8001", handler))
}
