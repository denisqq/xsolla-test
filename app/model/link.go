package model

import (
	"github.com/gofrs/uuid"
)

type Link struct {
	BaseModel
	OriginUrl  string `gorm:"column:origin_url;size:256;not null"`
	ShortUrl   string `gorm:"column:short_url;size:8;not null"`
	Link       string `gorm:"column:link;size:256;not null"`
	UserId     uuid.UUID
	Conversion int `gorm:"column:conversion;default:0"`
}
