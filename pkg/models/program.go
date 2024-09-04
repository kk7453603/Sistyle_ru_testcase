package models

import (
	"github.com/google/uuid"
)

type ProgramItem struct {
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	NameEn    string    `json:"nameEn"`
	IsPublic  bool      `json:"isPublic"`
	ProjectID uuid.UUID `json:"projectId"`
}

type ProgramList struct {
	Items []ProgramItem `json:"items"`
}
