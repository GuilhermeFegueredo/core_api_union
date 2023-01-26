package models

type User struct {
	User_ID  uint64 `json:"user_id,omitempty"`
	Email    string `json:"user_email,omitempty"`
	Password string `json:"user_pwd,omitempty"`
}
