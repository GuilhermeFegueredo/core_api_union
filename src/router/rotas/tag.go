package rotas

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
	// {
	// 	URI:             "/tags/{ID}",
	// 	Method:          http.MethodGet,
	// 	Function:        controllers.GetTag,
	// 	IsAuthenticated: false,
	// },
	{
		URI:             "/tags",
		Method:          http.MethodPost,
		Function:        controllers.CreateTag,
		IsAuthenticated: false,
	},
	// {
	// 	URI:             "/tags",
	// 	Method:          http.MethodPut,
	// 	Function:        controllers.UpdateTag,
	// 	IsAuthenticated: false,
	// },
	// {
	// 	URI:             "/tags",
	// 	Method:          http.MethodDelete,
	// 	Function:        controllers.DeleteTag,
	// 	IsAuthenticated: false,
	// },
}
