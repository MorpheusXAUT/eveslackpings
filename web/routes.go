package web

import (
	"net/http"
)

// Route stores information about a web route being handled
type Route struct {
	// Name represents a name for the web route
	Name string
	// Methods contains all HTTP methods available to access this route
	Methods []string
	// Pattern defines the URL pattern used to match this route
	Pattern string
	// HandlerFunc represents the web handler function to call for this route
	HandlerFunc http.HandlerFunc
}

// SetupRoutes initialises all used web routes and returns them for the router
func SetupRoutes(controller *Controller) []Route {
	r := []Route{
		Route{
			Name:        "PingPost",
			Methods:     []string{"POST"},
			Pattern:     "/ping",
			HandlerFunc: controller.PingPostHandler,
		},
	}

	return r
}
