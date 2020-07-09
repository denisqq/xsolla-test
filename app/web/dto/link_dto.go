package dto

import "github.com/gofrs/uuid"

type LinkDto struct {
	ID         uuid.UUID `json:"id"`
	Link       string    `json:"url"`
	Conversion int       `json:"conversion"`
}
