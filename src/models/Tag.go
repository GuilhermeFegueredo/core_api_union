package models

type Tag struct {
	Tag_ID   uint64  `json:"tag_id,omitempty"`
	Tag_Name string  `json:"tag_name,omitempty"`
	Tag_Type *uint64 `json:"tag_type,omitempty"`
}
