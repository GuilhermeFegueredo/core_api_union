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
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := db.Conectar()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryByTag(db)
	tag, err := repository.GetTag(ID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, tag)
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
		w.Write([]byte("Error connecting to database"))
		// Aqui entrará o sistema de respostas
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

func DeleteTag(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	ID, err := strconv.ParseUint(parameters["ID"], 10, 32)
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

	repository := repositories.NewRepositoryByTag(db)
	err = repository.DeleteTag(ID)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	response.JSON(w, http.StatusOK, map[string]string{"message": "Tag deleted successfully"})
}
