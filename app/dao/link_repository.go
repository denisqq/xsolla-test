package dao

import (
	"github.com/denisqq/xsolla-test/app/model"
	"github.com/gofrs/uuid"
	"github.com/jinzhu/gorm"
)

type LinkRepository struct {
	DB *gorm.DB
}

func NewLinkRepository(DB *gorm.DB) LinkRepository {
	return LinkRepository{DB: DB}
}

func (r *LinkRepository) Create(link model.Link) model.Link {
	r.DB.Create(&link)
	return link
}

func (r *LinkRepository) FindByLinkId(id uuid.UUID) (model.Link, error) {
	var link model.Link
	result := r.DB.Where(model.Link{BaseModel: model.BaseModel{ID: id}}).Find(&link)
	if result.Error != nil {
		return link, result.Error
	}
	return link, nil
}

func (r *LinkRepository) FindByShortUrl(shortUrl string) (model.Link, error) {
	var link model.Link
	result := r.DB.Where(model.Link{ShortUrl: shortUrl}).Find(&link)
	if result.Error != nil {
		return link, result.Error
	}
	return link, nil
}

func (r *LinkRepository) FindLinksByUser(username string) ([]model.Link, error) {
	var links []model.Link
	result := r.DB.Joins("JOIN users on users.id=links.user_id").Where("users.username = ?", username).Find(&links).Group("users.id")
	if result.Error != nil {
		return nil, result.Error
	}
	return links, nil
}

func (r *LinkRepository) DeleteLink(linkId uuid.UUID) error {
	 link := model.Link{BaseModel: model.BaseModel{ID: linkId}}
	 err := r.DB.Delete(&link).Error
	 if err != nil {
	 	return err
	 }

	 history := model.LinkHistory{LinkId: linkId}
	 err = r.DB.Delete(&history).Error
	 if err != nil {
	 	return err
	 }

	 return nil
}

func (r *LinkRepository) IncreaseLinkConversation(link model.Link) error {
	return r.DB.Transaction(func(tx *gorm.DB) error {
		r.DB.Where(link).Find(&link)
		conversion := link.Conversion + 1
		if err := tx.Model(&link).Update("conversion", conversion).Error; err != nil {
			tx.Rollback()
			return err
		}
		return nil
	})
}
