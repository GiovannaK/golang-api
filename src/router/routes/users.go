package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Router{
	{
		URI:    "/users",
		Method: http.MethodPost,
		Func: controllers.CreateUser,
		AuthRequired: false,
	},
	{
		URI:    "/users",
		Method: http.MethodGet,
		Func: controllers.SearchUsers,
		AuthRequired: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodGet,
		Func: controllers.SearchUser,
		AuthRequired: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodPut,
		Func: controllers.UpdateUser,
		AuthRequired: false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodDelete,
		Func: controllers.DeleteUser,
		AuthRequired: false,
	},
}	
