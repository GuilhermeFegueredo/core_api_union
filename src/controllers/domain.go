package controllers

import (
	"core_APIUnion/src/db"
	"core_APIUnion/src/repositories"
	"core_APIUnion/src/response"
	"errors"
	"net/http"
)

func GetDomainByName(w http.ResponseWriter, r *http.Request) {
	db, err := db.Conectar()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryByDomain(db)
	domain, err := repository.GetDomainByName(r.URL.Query().Get("name"))
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if domain == nil {
		response.Erro(w, http.StatusNotFound, errors.New("domain not found"))
		return

	}

	response.JSON(w, http.StatusOK, domain)
}
