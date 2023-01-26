package rotas

import (
	"core_APIUnion/src/controllers"
	"net/http"
)

var routerCostumer = []Routes{
	{
		URI:             "/costumer",
		Method:          http.MethodGet,
		Function:        controllers.GetCostumers,
		IsAuthenticated: false,
	},
	{
		URI:             "/costumer/name/{name}",
		Method:          http.MethodGet,
		Function:        controllers.GetCostumerByName,
		IsAuthenticated: false,
	},
	{
		URI:             "/costumer/id/{id}",
		Method:          http.MethodGet,
		Function:        controllers.GetCostumerByID,
		IsAuthenticated: false,
	},
	{
		URI:             "/costumer",
		Method:          http.MethodPost,
		Function:        controllers.CreateCostumer,
		IsAuthenticated: false,
	},
	{
		URI:             "/costumer/update/{id}",
		Method:          http.MethodPut,
		Function:        controllers.UpdateCostumer,
		IsAuthenticated: false,
	},
	{
		URI:             "/costumer/delete/{id}",
		Method:          http.MethodPut,
		Function:        controllers.DeleteCostumer,
		IsAuthenticated: false,
	},
}