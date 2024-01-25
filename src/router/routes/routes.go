package routes

import "net/http"

// Router represents all routes from API
type Router struct {
	URI   string
	Method string
	Func func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}