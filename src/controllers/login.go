package controllers

import (
	"core_APIUnion/src/auth"
	"core_APIUnion/src/db"
	"core_APIUnion/src/models"
	"core_APIUnion/src/repositories"
	"core_APIUnion/src/response"
	"core_APIUnion/src/security"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.Erro(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		response.Erro(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Conectar()
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repo := repositories.NewRepositoryByUser(db)
	userFromDB, err := repo.GetUserByEmail(user.Email)
	if err != nil {
		response.Erro(w, http.StatusInternalServerError, err)
		return
	}

	err = security.VerificarSenha(userFromDB.Password, user.Password)
	if err != nil {
		response.Erro(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := auth.NewToken(userFromDB.ID)
	response.JSON(w, http.StatusAccepted, token)
}
