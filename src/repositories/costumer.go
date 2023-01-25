package repositories

import (
	"core_APIUnion/src/models"
	"database/sql"
	"log"
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

		if err = lines.Scan(&costumer.Costumer_ID, &costumer.Costumer_name, &costumer.Status_ID); err != nil {
			log.Fatal("Error scanning costumer ", err) // Aqui entrará o sistema de respostas
			return nil, err
		}

		costumers = append(costumers, costumer)
	}

	return costumers, nil

}
