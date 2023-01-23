package router

import (
	"core_APIUnion/src/router/rotas"

	"github.com/gorilla/mux"
)

// Gerar vai retonrar um router com as rotas configuradas
func Gerar() *mux.Router {
	r := mux.NewRouter()
	return rotas.Configurar(r)
}
