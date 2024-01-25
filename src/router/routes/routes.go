package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Router represents all routes from API
type Router struct {
	URI   string
	Method string
	Func func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func Configure(r *mux.Router) *mux.Router {
	routes := routesUsers
	for _, route := range routes {
		r.HandleFunc(route.URI, route.Func).Methods(route.Method)
	}
	return r
}