package response

import (
	"encoding/json"
	"net/http"
)

// Eduardo Lisboa: Repostas para serem usadas nas controlers :D
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dados)
}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
