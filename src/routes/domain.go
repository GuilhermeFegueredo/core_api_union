package routes

import (
	"core_APIUnion/src/controllers"
	"net/http"
)

var routerDomain = []Routes{
	{
		URI:             "/domain",
		Method:          http.MethodGet,
		Function:        controllers.GetDomainByName,
		IsAuthenticated: false,
	},
	{
		URI:             "/domain",
		Method:          http.MethodPost,
		Function:        controllers.CreateDomain,
		IsAuthenticated: false,
	},
	{
		URI:             "/domain/{id}",
		Method:          http.MethodPut,
		Function:        controllers.UpdateDomain,
		IsAuthenticated: false,
	},
	{
		URI:             "/domain/{id}",
		Method:          http.MethodDelete,
		Function:        controllers.DeleteDomain,
		IsAuthenticated: false,
	},
}
