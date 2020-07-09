package dao

import (
	"github.com/denisqq/xsolla-test/app/model"
	"github.com/jinzhu/gorm"
)

type LinkHistoryRepository struct {
	DB *gorm.DB
}

func NewLinkHistoryRepository(DB *gorm.DB) LinkHistoryRepository {
	return LinkHistoryRepository{DB: DB}
}

func (r *LinkHistoryRepository) CreateLinkConversionHistory(history model.LinkHistory) {
	r.DB.Create(&history)
}
