package repositories

import (
	"core_APIUnion/src/models"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

func NewRepositoryByUser(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

func (repo Usuarios) GetUserByEmail(email string) (models.User, error) {
	linha, err := repo.db.Query("select user_id, user_pwd from tblUser where user_email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer linha.Close()

	var usuario models.User

	if linha.Next() {
		err = linha.Scan(&usuario.ID, &usuario.Password)
		if err != nil {
			return models.User{}, err
		}
	}

	return usuario, nil
}
