package dao

import (
	"fmt"
	"github.com/denisqq/xsolla-test/app/model"
	"github.com/denisqq/xsolla-test/app/web/dto"
	"github.com/denisqq/xsolla-test/app/web/request"
	"github.com/jinzhu/gorm"
)

type StatRepository struct {
	DB *gorm.DB
}

func NewStatRepository(DB *gorm.DB) StatRepository {
	return StatRepository{DB: DB}
}

const DayFormat = "%Y-%m-%d"
const HourFormat = "%H"
const MinuteFormat = "%i"

func (r *StatRepository) GetTop(limit int) ([]model.Link, error) {
	var links []model.Link
	if err := r.DB.Limit(limit).Order("conversion desc").Find(&links).Error; err != nil {
		return nil, err
	}
	return links, nil
}

func (r *StatRepository) ConversionGraph(by request.GroupBy) ([]dto.ConversionStat, error) {
	var format string

	//TODO Чет не так
	if by == request.Day {
		format = DayFormat
	}
	if by == request.Hour {
		format = DayFormat + " " + HourFormat
	}
	if by == request.Minute {
		format = DayFormat + " " + HourFormat + "-" + MinuteFormat
	}

	var conversions []dto.ConversionStat
	query := fmt.Sprintf("select date_format(lh.created_at,'%s') date, count(1) conversion "+
		"from link_histories lh "+
		"group by date "+
		"order by date desc", format)

	err := r.DB.Raw(query).Scan(&conversions).Error

	if err != nil {
		return nil, err
	}

	return conversions, nil
}
