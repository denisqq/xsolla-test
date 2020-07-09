package service

import (
	"github.com/denisqq/xsolla-test/app/dao"
	"github.com/denisqq/xsolla-test/app/model"
	"github.com/denisqq/xsolla-test/app/web/dto"
	"github.com/denisqq/xsolla-test/app/web/request"
)

type StatService struct {
	dao.StatRepository
}

func NewStatService(statRepository dao.StatRepository) StatService {
	return StatService{StatRepository: statRepository}
}

func (service *StatService) GetTop(limit int) ([]model.Link, error) {
	return service.StatRepository.GetTop(limit)
}

func (service *StatService) GetConversionGraph(req request.StatConversionGraphRequest) ([]dto.ConversionStat, error) {
	return service.StatRepository.ConversionGraph(req.GroupBy)
}