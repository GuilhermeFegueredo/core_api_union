package repositories

import (
	"core_APIUnion/src/models"
	"database/sql"
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
		return nil, err
	}

	defer lines.Close()

	var tags []models.Tag

	for lines.Next() {
		var tag models.Tag

		if err = lines.Scan(&tag.Tag_ID, &tag.Tag_Name, &tag.Domain_value); err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil

}

func (repository Tags) GetTag(ID uint64) (models.Tag, error) {

	stmt, err := repository.db.Prepare("SELECT T.tag_id, T.tag_name, D.domain_value FROM tblTags T INNER JOIN tblDomain D ON T.tag_type = D.domain_id WHERE T.tag_id = ?")
	if err != nil {
		return models.Tag{}, err
	}

	defer stmt.Close()
	var tag models.Tag

	err = stmt.QueryRow(ID).Scan(&tag.Tag_ID, &tag.Tag_Name, &tag.Domain_value)
	if err != nil {
		return models.Tag{}, err
	}

	return tag, nil
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

func (repository Tags) DeleteTag(ID uint64) error {
	stmt, err := repository.db.Prepare("DELETE FROM tblTags WHERE tag_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}
