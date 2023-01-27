package repositories

import (
	"core_APIUnion/src/models"
	"database/sql"
)

type Status struct {
	db *sql.DB
}

// NewRepositoryByStatus - constrói uma estrutura de db que usa um ponteiro sql.DB como argumento
func NewRepositoryByStatus(db *sql.DB) *Status {
	return &Status{db}
}

// GetStatusById - recebe um id uint64 como argumento e retorna uma estrutura models.Status
func (repository Status) GetStatusById(id uint64) (models.Status, error) {

	stmt, err := repository.db.Prepare("SELECT * FROM tblStatus WHERE status_id = ? ")
	if err != nil {
		return models.Status{}, err
	}
	defer stmt.Close()

	var status models.Status
	err = stmt.QueryRow(id).Scan(&status.Status_ID, &status.Dominio, &status.Order, &status.Description)

	if err != nil {
		return models.Status{}, err
	}

	return status, nil
}
