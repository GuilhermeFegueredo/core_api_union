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

// GetDomains puxa todos os domains por nome requerido
func GetDomainByName(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
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

// CreateDomain cria um novo domain
func CreateDomain(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var domain models.Domain
	if err = json.Unmarshal(bodyRequest, &domain); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = domain.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryByDomain(db)
	domain.Domain_ID, err = repository.CreateDomain(domain)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, domain)
}

// UpdateDomain altera as informações de um domain no banco
func UpdateDomain(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	domain_ID, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var domain models.Domain
	if err = json.Unmarshal(bodyRequest, &domain); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = domain.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryByDomain(db)
	err = repository.UpdateDomain(domain_ID, domain)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, domain)

}

// DeleteDomain deleta um domain do banco
func DeleteDomain(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)
	domain_ID, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.NewRepositoryByDomain(db)
	if err = repository.DeleteDomain(domain_ID); err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSONMessage(w, http.StatusOK, "domain successfully deleted")
}
