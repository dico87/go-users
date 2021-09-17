package domain

import (
	"time"
)

type User struct {
	ID             uint         `json:"id"`
	DocumentTypeID uint         `json:"document_type_id"`
	DocumentType   DocumentType `json:"document_type"`
	Document       string       `json:"document"`
	LastName       string       `json:"last_name"`
	SurName        string       `json:"sur_name"`
	Name           string       `json:"name"`
	OtherNames     string       `json:"other_names"`
	Birthday       time.Time    `json:"birthday"`
	Sex            string       `json:"sex"`
	Active         bool         `json:"active"`
	Photo          string       `json:"photo"`
}

type DocumentType struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
}
