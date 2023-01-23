package rotas

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
}
