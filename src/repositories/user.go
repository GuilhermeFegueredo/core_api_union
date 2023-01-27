package repositories

import (
	"core_APIUnion/src/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

// NewRepositoryByUser cria um novo repositorio de usuarios
func NewRepositoryByUser(db *sql.DB) *Users {
	return &Users{db}
}

// CreateUser cria um novo usuario no banco de dados
func (repository Users) CreateUser(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare("INSERT INTO tblUser(user_name, user_email,user_level,user_pwd, status_id) VALUES(?,?,?,?,?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Email, user.Level, user.Password, user.Status_ID)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

// GetUsers retorna todos os usuarios do banco de dados
func (repository Users) GetUsers(owner string) ([]models.User, error) {
	owner = "%" + owner + "%"
	rows, err := repository.db.Query("SELECT user_id, user_name, user_email,user_level, status_id FROM tblUser WHERE user_name LIKE ? OR user_email LIKE ?", owner, owner)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.User_ID, &user.Name, &user.Email, &user.Level, &user.Status_ID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

// GetUser retorna um usuario do banco de dados pelo id
func (repository Users) GetUser(id uint64) (models.User, error) {
	row := repository.db.QueryRow("SELECT user_id, user_name, user_email,user_level, status_id FROM tblUser WHERE user_id = ?", id)

	var user models.User
	if err := row.Scan(&user.User_ID, &user.Name, &user.Email, &user.Level, &user.Status_ID); err != nil {
		return user, err
	}

	return user, nil
}

// UpdateUser atualiza um usuario no banco de dados pelo id
func (repository Users) UpdateUser(id uint64, user models.User) error {
	statement, err := repository.db.Prepare("UPDATE tblUser SET user_name = ?, user_email = ?,user_level = ?, status_id = ? WHERE user_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(user.Name, user.Email, user.Level, user.Status_ID, id); err != nil {
		return err
	}

	return nil
}

// DeleteUser deleta um usuario do banco de dados pelo id
func (repository Users) DeleteUser(id uint64) error {
	statement, err := repository.db.Prepare("DELETE FROM tblUser WHERE user_id = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(id); err != nil {
		return err
	}

	return nil
}

// GetUserByEmail retorna um usuario do banco de dados pelo email
func (repo Users) GetUserByEmail(email string) (models.User, error) {
	linha, err := repo.db.Query("select user_id, user_pwd from tblUser where user_email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer linha.Close()

	var usuario models.User

	if linha.Next() {
		err = linha.Scan(&usuario.User_ID, &usuario.Password)
		if err != nil {
			return models.User{}, err
		}
	}

	return usuario, nil
}
