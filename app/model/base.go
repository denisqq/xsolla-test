package model

import (
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
	"time"
)

type BaseModel struct {
	ID        uuid.UUID `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (base *BaseModel) BeforeCreate(scope *gorm.Scope) error {
	v4, _ := uuid.NewV4()
	_ = scope.SetColumn("CreatedAt", time.Now())
	_ = scope.SetColumn("UpdatedAt", time.Now())
	return scope.SetColumn("ID", v4)
}

func (base *BaseModel) BeforeUpdate(scope *gorm.Scope) error {
	return scope.SetColumn("UpdatedAt", time.Now())
}
