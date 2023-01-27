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

// GetCostumers - lista todos os customers ativos
func GetCostumers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryByCostumer(db)
	costumers, err := repository.GetCostumers()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	err = json.NewEncoder(w).Encode(costumers)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}
}

// GetCostumerByName - lista os customers ativos pelo nome
func GetCostumerByName(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	costumerName := parametros["name"]

	db, err := db.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryByCostumer(db)
	costumer, err := repository.GetCostumerByName(costumerName)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if costumer == nil {
		response.Erro(w, http.StatusNotFound, errors.New("costumer not found"))
		return
	}

	response.JSON(w, http.StatusOK, costumer)
}

// GetCostumerByID - lista os customers ativos pelo ID
func GetCostumerByID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	costumer_id, err := strconv.ParseUint(parameters["id"], 10, 64)
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

	repository := repositories.NewRepositoryByCostumer(db)
	costumer, err := repository.GetCostumerByID(costumer_id)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, costumer)
}

// CreateCostumer - Cria o customer 
func CreateCostumer(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var costumer models.Costumer
	if err = json.Unmarshal(bodyRequest, &costumer); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}


	if err = costumer.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)

	costumer.Status_ID = 3

		return
	}

	db, err := db.Connect()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryByCostumer(db)
	costumer.Costumer_ID, err = repository.CreateCostumer(costumer)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	costumer, err = repository.GetCostumerByID(costumer.Costumer_ID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, costumer)
}

// UpdateCostumer - Altera dados do customer
func UpdateCostumer(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	costumer_id, err := strconv.ParseUint(parametros["id"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var costumer models.Costumer
	if err = json.Unmarshal(bodyRequest, &costumer); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	if err = costumer.Prepare(); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, erro := db.Connect()
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryByCostumer(db)
	costumer, err = repository.UpdateCostumer(costumer_id, costumer)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, costumer)
}

// DeleteCostumer - soft delete no customer
func DeleteCostumer(w http.ResponseWriter, r *http.Request) {
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

	repository := repositories.NewRepositoryByCostumer(db)
	costumer, err := repository.DeleteCostumer(id)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, costumer)
}
