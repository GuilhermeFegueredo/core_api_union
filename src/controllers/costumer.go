package controllers

import (
	"core_APIUnion/src/db"
	"core_APIUnion/src/models"
	"core_APIUnion/src/repositories"
	"core_APIUnion/src/response"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetCostumers(w http.ResponseWriter, r *http.Request) {
	db, err := db.Conectar()
	if err != nil {
		log.Fatal("Error connecting to database") // Aqui entrará o sistema de respostas
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryByCostumer(db)
	costumers, err := repository.GetCostumers()
	if err != nil {
		log.Fatal("Error fetching costumers") // Aqui entrará o sistema de respostas
		return
	}

	err = json.NewEncoder(w).Encode(costumers)
	if err != nil {
		log.Fatal("Error convert json") // Aqui entrará o sistema de respostas
		return
	}
}
func GetCostumerByName(w http.ResponseWriter, r *http.Request) {
	db, err := db.Conectar()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryByCostumer(db)
	costumer, err := repository.GetCostumerByName(r.URL.Query().Get("name"))
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	if costumer == nil {
		response.Erro(w, http.StatusNotFound, errors.New("Costumer not found"))
		return

	}

	response.JSON(w, http.StatusOK, costumer)
}

func GetCostumerByID(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	costumer_id, err := strconv.ParseUint(parameters["id"], 10, 64)
	if err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conectar()
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

func CreateCostumer(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		response.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var costumer models.Costumer
	if erro = json.Unmarshal(bodyRequest, &costumer); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = costumer.Prepare(); erro != nil {
		response.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, err := db.Conectar()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryByCostumer(db)
	costumer.Costumer_ID, erro = repository.CreateCostumer(costumer)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, costumer)
}
