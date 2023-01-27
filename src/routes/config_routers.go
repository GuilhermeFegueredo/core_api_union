package routes

import (
	"github.com/gorilla/mux"
)

// Generate - retorna um router com as rotas configuradas
func Generate() *mux.Router {
	r := mux.NewRouter()
	return Config(r)
}
