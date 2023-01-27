package controllers

import (
	"core_APIUnion/src/db"
	"core_APIUnion/src/repositories"
	"core_APIUnion/src/response"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetStatusById é uma função que recupera um status por seu ID do banco de dados.
func GetStatusById(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryByStatus(db)
	status, err := repository.GetStatusById(id)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, status)
}
