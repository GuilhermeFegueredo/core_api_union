package controllers

import (
	"core_APIUnion/src/db"
	"core_APIUnion/src/models"
	"core_APIUnion/src/repositories"
	"core_APIUnion/src/response"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
func CreateCostumers(w http.ResponseWriter, r *http.Request) {
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
	costumer.Costumer_ID, erro = repository.CreateCostumers(costumer)
	if erro != nil {
		response.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	response.JSON(w, http.StatusOK, costumer)
}
