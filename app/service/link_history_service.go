package service

import (
	"github.com/denisqq/xsolla-test/app/dao"
	"github.com/denisqq/xsolla-test/app/model"
	"github.com/gin-gonic/gin"
)

type LinkHistoryService struct {
	LinkHistoryRepository dao.LinkHistoryRepository
	LinkService           LinkService
}

func NewLinkHistoryService(linkHistoryRepository dao.LinkHistoryRepository, linkService LinkService) LinkHistoryService {
	return LinkHistoryService{LinkHistoryRepository: linkHistoryRepository, LinkService: linkService}
}

func (service *LinkHistoryService) WriteHistoryLinkConversion(c *gin.Context) {
	go func() {
		shortUrl := c.Param("shortUrl")

		link, _ := service.LinkService.GetShortLink(shortUrl)
		history := model.LinkHistory{
			LinkId: link.ID,
		}
		service.LinkHistoryRepository.CreateLinkConversionHistory(history)
		service.LinkService.IncreaseLinkConversation(link)
	}()
}
