package server

import (
	"net/http"

	"github.com/DudeWhoCode/yellowjacket/backend"
)

var b = backend.GetBroker()

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

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
		AttackHandler,
	},
	Route{
		"Home",
		"GET",
		"/",
		MainHandler,
	},
	Route{
		"Events",
		"GET",
		"/events",
		b.ServeHTTP,
	},
	// Route{
	// 	"TableDescribe",
	// 	"GET",
	// 	"/api/v1/describe-table/{schema}/{name}",
	// 	TableDesc,
	// },
}
