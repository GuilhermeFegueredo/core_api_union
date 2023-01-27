package routes

import (
	"core_APIUnion/src/controllers"
	"net/http"
)

var routerTags = []Routes{
	{
		URI:             "/tags",
		Method:          http.MethodGet,
		Function:        controllers.GetTags,
		IsAuthenticated: false,
	},
	{
		URI:             "/tag/{id}",
		Method:          http.MethodGet,
		Function:        controllers.GetTag,
		IsAuthenticated: false,
	},
	{
		URI:             "/tag",
		Method:          http.MethodPost,
		Function:        controllers.CreateTag,
		IsAuthenticated: false,
	},
	// {
	// 	URI:             "/tag/update/{id}",
	// 	Method:          http.MethodPut,
	// 	Function:        controllers.UpdateTag,
	// 	IsAuthenticated: false,
	// },
	{
		URI:             "/tag/delete/{id}",
		Method:          http.MethodDelete,
		Function:        controllers.DeleteTag,
		IsAuthenticated: false,
	},
}
