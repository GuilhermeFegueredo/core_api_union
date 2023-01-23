package repositories

import (
	"core_APIUnion/src/models"
	"database/sql"
	"log"
)

type Tags struct {
	db *sql.DB
}

func NewRepositoryByTag(db *sql.DB) *Tags {
	return &Tags{db}
}

func (repository Tags) GetTags() ([]models.Tag, error) {
	lines, err := repository.db.Query("SELECT * FROM tblTags")
	if err != nil {
		log.Fatal("Error selecting tags") // Aqui entrará o sistema de respostas
		return nil, err
	}

	defer lines.Close()

	var tags []models.Tag

	for lines.Next() {
		var tag models.Tag

		if err = lines.Scan(&tag.Tag_ID, &tag.Tag_Name, &tag.Tag_Type); err != nil {
			log.Fatal("Error scanning tag ", err) // Aqui entrará o sistema de respostas
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil

}
