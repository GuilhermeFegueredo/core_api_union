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

func CreateTag(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
	}

	var tag models.Tag
	if err = json.Unmarshal(bodyRequest, &tag); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
	}

	db, err := db.Conectar()
	if err != nil {
		log.Fatal("Erro connecting to database") // Aqui entrará o sistema de respostas
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryByTag(db)

	tag.Tag_ID, err = repository.CreateTag(tag)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
	}

	response.JSON(w, http.StatusCreated, tag)

}
