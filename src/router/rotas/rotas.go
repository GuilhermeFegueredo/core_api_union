package rotas

import (
	"net/http"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da API
type Routes struct {
	URI             string
	Method          string
	Function        func(http.ResponseWriter, *http.Request)
	IsAuthenticated bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(r *mux.Router) *mux.Router {
	routers := routerTags
	routers = append(routers, routerDomain...)
	routers = append(routers, routerStatus...)
	routers = append(routers, RouterUser...)

	for _, router := range routers {
		r.HandleFunc(router.URI, router.Function).Methods(router.Method)
	}

	return r
}
