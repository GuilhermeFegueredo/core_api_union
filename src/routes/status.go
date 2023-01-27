package routes

import (
	"core_APIUnion/src/controllers"
	"net/http"
)

var routerStatus = []Routes{
	{
		URI:             "/status/{id}",
		Method:          http.MethodGet,
		Function:        controllers.GetStatusById,
		IsAuthenticated: false,
	},
}
