package models

type Status struct {
	Status_ID   uint64 `json:"status_id,omitempty"`
	Dominio     string `json:"dominio,omitempty"`
	Order       uint64 `json:"order,omitempty"`
	Description string `json:"description,omitempty"`
}
