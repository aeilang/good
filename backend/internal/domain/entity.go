package domain

import "github.com/google/uuid"

type IndexPost struct {
	ID          uuid.UUID
	Title       string
	Description string
	Requirement string
}
