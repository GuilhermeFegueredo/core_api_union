package rotas

import (
	"core_APIUnion/src/controllers"
	"net/http"
)

var routerLogin = []Routes{
	{
		URI:             "/login",
		Method:          http.MethodPost,
		Function:        controllers.Login,
		IsAuthenticated: false,
	},
}
