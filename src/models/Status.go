package models

type Status struct {
	Status_ID   uint64 `json:"status_id"`
	Dominio     string `json:"dominio"`
	Order       uint64 `json:"order"`
	Description string `json:"description"`
}
