package controllers

import (
	"core_APIUnion/src/db"
	"core_APIUnion/src/repositories"
	"core_APIUnion/src/response"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func GetTag(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["ID"], 10, 32)
	if err != nil {
		w.Write([]byte("Error to convert parameter to int"))
		return
	}

	db, err := db.Conectar()
	if err != nil {
		w.Write([]byte("Error connecting to database"))
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryByTag(db)
	tag, err := repository.GetTag(ID)
	if err != nil {
		w.Write([]byte("Error"))
	}

	response.JSON(w, http.StatusOK, tag)
}
