package server

import (
	"net/http"

	"github.com/DudeWhoCode/yellowjacket/backend"
)

var b = backend.GetBroker()

// Route struct is used to create all routes the webserver uses
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a collection of individial route/url
type Routes []Route

var routes = Routes{
	Route{
		"ping",
		"GET",
		"/ping",
		Ping,
	},
	Route{
		"Attack",
		"POST",
		"/api/v1/attack",
		StartSwarm,
	},
	Route{
		"Home",
		"GET",
		"/",
		Home,
	},
	Route{
		"Events",
		"GET",
		"/events",
		b.ServeHTTP,
	},
}
