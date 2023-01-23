package models

type Domain struct {
	Domain_ID    uint64 `json:"domain_id,omitempty"`
	Domain_name  string `json:"domain_name,omitempty"`
	Domain_code  uint64 `json:"domain_code,omitempty"`
	Domain_value string `json:"domain_value,omitempty"`
}


