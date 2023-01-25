package controllers

import (
	"core_APIUnion/src/db"
	"core_APIUnion/src/repositories"
	"encoding/json"
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
