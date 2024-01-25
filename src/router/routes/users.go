package routes

import "net/http"

var routesUsers = []Router{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodGet,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodPut,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodDelete,
		Func: func(w http.ResponseWriter, r *http.Request) {

		},
		AuthRequired: false,
	},
}	
