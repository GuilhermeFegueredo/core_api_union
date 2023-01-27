package repositories

import (
	"core_APIUnion/src/models"
	"database/sql"
	"fmt"
	"log"
)

type Costumers struct {
	db *sql.DB
}

func NewRepositoryByCostumer(db *sql.DB) *Costumers {
	return &Costumers{db}
}

func (repository Costumers) GetCostumers() ([]models.Costumer, error) {
	lines, err := repository.db.Query("SELECT C.costumer_id, C.costumer_name, S.status_description FROM tblCostumer C INNER JOIN tblStatus S ON C.status_id = S.status_id WHERE C.status_id = 3")
	if err != nil {
		log.Fatal("Error selecting costumers") // Aqui entrará o sistema de respostas
		return nil, err
	}

	defer lines.Close()

	var costumers []models.Costumer

	for lines.Next() {
		var costumer models.Costumer

		if err = lines.Scan(&costumer.Costumer_ID, &costumer.Costumer_name, &costumer.Description); err != nil {
			log.Fatal("Error scanning costumer ", err) // Aqui entrará o sistema de respostas
			return nil, err
		}

		costumers = append(costumers, costumer)
	}

	return costumers, nil
}

func (repository Costumers) GetCostumerByName(name string) ([]models.Costumer, error) {
	text := "%' AND C.status_id = 3"
	query := fmt.Sprint("SELECT C.costumer_id, C.costumer_name, S.status_description FROM tblCostumer C INNER JOIN tblStatus S ON C.status_id = S.status_id WHERE C.costumer_name LIKE '%", name, text)

	lines, err := repository.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var costumers []models.Costumer

	for lines.Next() {
		var costumer models.Costumer

		if err = lines.Scan(
			&costumer.Costumer_ID,
			&costumer.Costumer_name,
			&costumer.Description,
		); err != nil {
			return nil, err
		}

		costumers = append(costumers, costumer)
	}

	return costumers, nil
}

func (repository Costumers) GetCostumerByID(id uint64) (models.Costumer, error) {

	stmt, err := repository.db.Prepare("SELECT C.costumer_id, C.costumer_name, S.status_description FROM tblCostumer C INNER JOIN tblStatus S ON C.status_id = S.status_id WHERE costumer_id = ? AND C.status_id = 3")
	if err != nil {
		return models.Costumer{}, err
	}
	defer stmt.Close()

	var costumer models.Costumer
	err = stmt.QueryRow(id).Scan(&costumer.Costumer_ID, &costumer.Costumer_name, &costumer.Description)

	if err != nil {
		return models.Costumer{}, err
	}

	return costumer, nil
}

func (repository Costumers) CreateCostumer(costumer models.Costumer) (uint64, error) {
	statement, erro := repository.db.Prepare(
		"INSERT INTO tblCostumer (costumer_name, status_id) VALUES (?, 3)",
	)
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	result, erro := statement.Exec(costumer.Costumer_name)
	if erro != nil {
		return 0, erro
	}

	LastInsertId, erro := result.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	return uint64(LastInsertId), nil
}

func (repository Costumers) UpdateCostumer(id uint64, costumer models.Costumer) (models.Costumer, error) {
	stmt, err := repository.db.Prepare(
		"UPDATE tblCostumer SET costumer_name = ?, status_id = ? WHERE costumer_id = ?")
	if err != nil {
		return models.Costumer{}, err
	}

	if _, err = stmt.Exec(costumer.Costumer_name, costumer.Status_ID, id); err != nil {
		return models.Costumer{}, err
	}

	stmt, err = repository.db.Prepare("SELECT C.costumer_id, C.costumer_name, S.status_description FROM tblCostumer C INNER JOIN tblStatus S ON C.status_id = S.status_id WHERE costumer_id = ?")
	if err != nil {
		return models.Costumer{}, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&costumer.Costumer_ID, &costumer.Costumer_name, &costumer.Description)

	if err != nil {
		return models.Costumer{}, err
	}

	return costumer, nil
}

func (repository Costumers) DeleteCostumer(id uint64) (models.Costumer, error) {

	stmt, err := repository.db.Prepare(
		"UPDATE tblCostumer SET status_id = 4 WHERE costumer_id = ?")
	if err != nil {
		return models.Costumer{}, err
	}

	if _, err = stmt.Exec(id); err != nil {
		return models.Costumer{}, err
	}

	stmt, err = repository.db.Prepare("SELECT C.costumer_id, C.costumer_name, S.status_description FROM tblCostumer C INNER JOIN tblStatus S ON C.status_id = S.status_id WHERE costumer_id = ?")
	if err != nil {
		return models.Costumer{}, err
	}
	defer stmt.Close()

	var costumer models.Costumer
	err = stmt.QueryRow(id).Scan(&costumer.Costumer_ID, &costumer.Costumer_name, &costumer.Description)

	if err != nil {
		return models.Costumer{}, err
	}

	return costumer, nil
}
