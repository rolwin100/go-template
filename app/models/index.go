package models

type Index struct {
	IndexName string `json:"index_name" validate:"required"`
}
