package controllers

import (
	"core_APIUnion/src/db"
	"core_APIUnion/src/repositories"
	"encoding/json"
	"log"
	"net/http"
)

func GetTags(w http.ResponseWriter, r *http.Request) {
	db, err := db.Conectar()
	if err != nil {
		log.Fatal("Error connecting to database") // Aqui entrará o sistema de respostas
		return
	}

	defer db.Close()

	repository := repositories.NewRepositoryByTag(db)
	tags, err := repository.GetTags()
	if err != nil {
		log.Fatal("Error fetching tags") // Aqui entrará o sistema de respostas
		return
	}

	err = json.NewEncoder(w).Encode(tags)
	if err != nil {
		log.Fatal("Error convert json") // Aqui entrará o sistema de respostas
		return
	}
}
