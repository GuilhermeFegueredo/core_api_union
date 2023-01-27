package routes

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
func Config(r *mux.Router) *mux.Router {
	routes := routerTags
	routes = append(routes, routerDomain...)
	routes = append(routes, routerStatus...)
	routes = append(routes, routerLogin...)
	routes = append(routes, routerCostumer...)
	routes = append(routes, routerUser...)

	for _, router := range routes {
		r.HandleFunc(router.URI, router.Function).Methods(router.Method)
	}

	return r
}
