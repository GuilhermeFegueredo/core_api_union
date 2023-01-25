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
	lines, err := repository.db.Query("SELECT T.tag_id, T.tag_name, D.domain_value FROM tblTags T INNER JOIN tblDomain D ON T.tag_type = D.domain_id")
	if err != nil {
		log.Fatal("Error selecting tags") // Aqui entrará o sistema de respostas
		return nil, err
	}

	defer lines.Close()

	var tags []models.Tag

	for lines.Next() {
		var tag models.Tag

		if err = lines.Scan(&tag.Tag_ID, &tag.Tag_Name, &tag.Domain_value); err != nil {
			log.Fatal("Error scanning tag ", err) // Aqui entrará o sistema de respostas
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil

}

func (repository Tags) GetTag(ID uint64) (models.Tag, error) {

	stmt, err := repository.db.Prepare("SELECT T.tag_id, T.tag_name, D.domain_value FROM tblTags T INNER JOIN tblDomain D ON T.tag_type = D.domain_id WHERE T.tag_id = ?")
	if err != nil {
		log.Fatal("Error fetching tags")
		return models.Tag{}, err
	}

	defer stmt.Close()
	var tag models.Tag

	err = stmt.QueryRow(ID).Scan(&tag.Tag_ID, &tag.Tag_Name, &tag.Domain_value)
	if err != nil {
		log.Fatal("Error scan row")
		return models.Tag{}, err
	}

	return tag, nil
}
