package models

type Costumer struct {
	Costumer_ID   uint64 `json:"costumer_id,omitempty"`
	Costumer_name string `json:"costumer_name,omitempty"`
	Status        `json:"status_id"`
}
