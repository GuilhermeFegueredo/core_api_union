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

func (repository Tags) CreateTag(tag models.Tag) (uint64, error) {
	stmt, err := repository.db.Prepare("INSERT INTO tblTags (tag_name, tag_type) VALUES(?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	tagResp, err := stmt.Exec(tag.Tag_Name, tag.Domain_ID)
	if err != nil {
		return 0, err
	}

	LastInsertId, err := tagResp.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(LastInsertId), nil
}
