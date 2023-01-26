package models

import (
	"errors"
	"strings"
)

type Domain struct {
	Domain_ID    uint64 `json:"domain_id,omitempty"`
	Domain_name  string `json:"domain_name,omitempty"`
	Domain_code  uint64 `json:"domain_code,omitempty"`
	Domain_value string `json:"domain_value,omitempty"`
}

func (domain *Domain) format() error {
	domain.Domain_name = strings.TrimSpace(domain.Domain_name)
	domain.Domain_name = strings.ToUpper(domain.Domain_name)
	domain.Domain_value = strings.TrimSpace(domain.Domain_value)

	return nil
}

func (domain *Domain) validate() error {
	if domain.Domain_name == "" {
		return errors.New("the name is mandatory and cannot be blank")
	}

	if domain.Domain_value == "" {
		return errors.New("the value is mandatory and cannot be blank")
	}

	return nil
}

// Prepare vai chamar os métodos para validar e formatar o domain recebido
func (domain *Domain) Prepare() error {
	if erro := domain.validate(); erro != nil {
		return erro
	}

	if erro := domain.format(); erro != nil {
		return erro
	}

	return nil
}
