package models

type Tag struct {
	Tag_ID   uint64 `json:"tag_id,omitempty"`
	Tag_Name string `json:"tag_name,omitempty"`
	Domain   `json:"tag_type"`
}
