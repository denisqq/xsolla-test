package dto

import "github.com/gofrs/uuid"

type UserDto struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}
