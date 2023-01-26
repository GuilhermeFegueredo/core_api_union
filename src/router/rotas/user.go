package rotas

import (
	"core_APIUnion/src/controllers"
	"net/http"
)

var RouterUser = []Routes{
	{
		URI:             "/users",
		Method:          http.MethodGet,
		Function:        controllers.GetUsers,
		IsAuthenticated: false,
	},
	{
		URI:             "/users/{ID}",
		Method:          http.MethodGet,
		Function:        controllers.GetUser,
		IsAuthenticated: false,
	},
	{
		URI:             "/users",
		Method:          http.MethodPost,
		Function:        controllers.CreateUser,
		IsAuthenticated: false,
	},
	{
		URI:             "/users/{ID}",
		Method:          http.MethodPut,
		Function:        controllers.UpdateUser,
		IsAuthenticated: false,
	},
	{
		URI:             "/users/{ID}",
		Method:          http.MethodDelete,
		Function:        controllers.DeleteUser,
		IsAuthenticated: false,
	},
}
