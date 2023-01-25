package models

import (
	"errors"
	"strings"
)

type Tag struct {
	Tag_ID   uint64 `json:"tag_id,omitempty"`
	Tag_Name string `json:"tag_name,omitempty"`
	Domain   `json:"tag_type"`
}

func (tag *Tag) format() error {
	tag.Tag_Name = strings.TrimSpace(tag.Tag_Name)

	return nil
}

func (tag *Tag) validate() error {
	if tag.Tag_Name == "" {
		return errors.New("the name is mandatory and cannot be blank")
	}
	if tag.Domain_ID == 0 {
		return errors.New("the tag type is mandatory and cannot be blank")
	}

	return nil
}
func (tag *Tag) Prepare() error {
	if erro := tag.validate(); erro != nil {
		return erro
	}

	if erro := tag.format(); erro != nil {
		return erro
	}

	return nil
}
