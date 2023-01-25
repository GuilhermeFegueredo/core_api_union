package controllers

import (
	"core_APIUnion/src/db"
	"core_APIUnion/src/models"
	"core_APIUnion/src/repositories"
	"core_APIUnion/src/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func CreateDomain(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var domain models.Domain
	if erro = json.Unmarshal(bodyRequest, &domain); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = domain.Prepare(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, err := db.Conectar()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryByDomain(db)
	domain.Domain_ID, erro = repository.CreateDomain(domain)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, domain)
}

// UpdateDomain altera as informações de um domain no banco
func UpdateDomain(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	domain_ID, erro := strconv.ParseUint(parametros["domain_id"], 10, 64)
	if erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var domain models.Domain
	if erro = json.Unmarshal(bodyRequest, &domain); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = domain.Prepare(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := db.Conectar()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryByDomain(db)
	erro = repository.UpdateDomain(domain_ID, domain)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, domain)

}
