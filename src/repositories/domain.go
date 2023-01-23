package repositories

import (
	"core_APIUnion/src/models"
	"database/sql"
	"fmt"
	"strings"
)

type Domain struct {
	db *sql.DB
}

func NewRepositoryByDomain(db *sql.DB) *Domain {
	return &Domain{db}
}

func (repository Domain) GetDomainByName(name string) ([]models.Domain, error) {
	Domainame := strings.ToUpper(fmt.Sprintf("%%%s%%", name))

	lines, err := repository.db.Query("SELECT domain_value FROM tblDomain WHERE domain_name LIKE ? ", Domainame)
	if err != nil {
		return nil, err
	}

	defer lines.Close()

	var domains []models.Domain

	for lines.Next() {
		var domain models.Domain

		if err = lines.Scan(&domain.Domain_value); err != nil {
			return nil, err
		}

		domains = append(domains, domain)
	}

	return domains, nil
}
