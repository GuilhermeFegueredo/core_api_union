package repositories

import (
	"core_APIUnion/src/models"
	"database/sql"
	"fmt"
	"log"
	"strings"
)

type Costumers struct {
	db *sql.DB
}

func NewRepositoryByCostumer(db *sql.DB) *Costumers {
	return &Costumers{db}
}

func (repository Costumers) GetCostumers() ([]models.Costumer, error) {
	lines, err := repository.db.Query("SELECT C.costumer_id, C.costumer_name, S.status_description FROM tblCostumer C INNER JOIN tblStatus S ON C.status_id = S.status_id")
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
	Costumername := strings.ToUpper(fmt.Sprintf("%%%s%%", name))

	lines, err := repository.db.Query("SELECT costumer_id, costumer_name FROM tblCostumer WHERE costumer_name LIKE ?", Costumername)
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var costumers []models.Costumer

	for lines.Next() {
		var costumer models.Costumer

		if err = lines.Scan(&costumer.Costumer_ID, &costumer.Costumer_name); err != nil {
			return nil, err
		}

		costumers = append(costumers, costumer)
	}

	return costumers, nil
}

func (repository Costumers) GetCostumerByID(id uint64) (models.Costumer, error) {

	stmt, err := repository.db.Prepare("SELECT * FROM tblCostumer WHERE costumer_id = ? ")
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
		"insert into tblCostumer (costumer_name, status_id) values(?, 1)",
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
