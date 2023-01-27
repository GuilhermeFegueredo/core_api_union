package models

import (
	"core_APIUnion/src/security"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	User_ID    uint64    `json:"user_id,omitempty"`
	Name       string    `json:"user_name,omitempty"`
	Email      string    `json:"user_email,omitempty"`
	Level      int       `json:"user_level,omitempty"`
	Password   string    `json:"user_pwd,omitempty"`
	Created_At time.Time `json:"created_at,omitempty"`
	Status     `json:"status_id,omitempty"`
}

// Validate and format user data
func (user *User) Prepare(stage string) error {
	if error := user.validate(stage); error != nil {
		return error
	}

	if error := user.format(stage); error != nil {
		return error
	}

	return nil

}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("required name")
	}

	if user.Email == "" {
		return errors.New("required email")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("invalid email")
	}

	if user.Level == 0 {
		return errors.New("required level")
	}

	if user.Password == "" {
		return errors.New("required password")
	}
	if user.Status_ID == 0 {
		return errors.New("required status")
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {

		passwordHash, err := security.HashPassword(user.Password)
		if err != nil {
			return errors.New("error to encrypt password")

		}

		user.Password = string(passwordHash)
	}

	return nil
}
