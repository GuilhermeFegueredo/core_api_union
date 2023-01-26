package models

import (
	"errors"
	"strings"
)

type Costumer struct {
	Costumer_ID   uint64 `json:"costumer_id,omitempty"`
	Costumer_name string `json:"costumer_name,omitempty"`
	Status        `json:"status_id"`
}

func (costumer *Costumer) Prepare() error {
	if erro := costumer.validate(); erro != nil {
		return erro
	}

	if erro := costumer.format(); erro != nil {
		return erro
	}

	return nil
}

func (costumer *Costumer) format() error {
	costumer.Costumer_name = strings.TrimSpace(costumer.Costumer_name)

	return nil
}

func (costumer *Costumer) validate() error {
	if costumer.Costumer_name == "" {
		return errors.New("The name is mandatory and cannot be blank")
	}

	return nil
}
