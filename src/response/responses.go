package response

import (
	"encoding/json"
	"log"
	"net/http"
)

// Eduardo Lisboa: Repostas para serem usadas nas controlers :D
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")

	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			log.Fatal(erro)
		}
	}
}

func Erro(w http.ResponseWriter, statusCode int, erro error) {
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}

func JSONMessage(w http.ResponseWriter, statusCode int, message string) {
	JSON(w, statusCode, struct {
		Message string `json:"message"`
	}{
		Message: message,
	})
}
