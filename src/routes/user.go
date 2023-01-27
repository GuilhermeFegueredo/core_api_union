package routes

import (
	"core_APIUnion/src/controllers"
	"net/http"
)

var routerUser = []Routes{
	{
		URI:             "/users",
		Method:          http.MethodGet,
		Function:        controllers.GetUsers,
		IsAuthenticated: false,
	},
	{
		URI:             "/user/{id}",
		Method:          http.MethodGet,
		Function:        controllers.GetUser,
		IsAuthenticated: false,
	},
	{
		URI:             "/user",
		Method:          http.MethodPost,
		Function:        controllers.CreateUser,
		IsAuthenticated: false,
	},
	{
		URI:             "/user/update/{id}",
		Method:          http.MethodPut,
		Function:        controllers.UpdateUser,
		IsAuthenticated: false,
	},
	{
		URI:             "/user/delete/{id}",
		Method:          http.MethodDelete,
		Function:        controllers.DeleteUser,
		IsAuthenticated: false,
	},
}
