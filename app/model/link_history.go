package model

import "github.com/gofrs/uuid"

type LinkHistory struct {
	BaseModel
	LinkId uuid.UUID
}
