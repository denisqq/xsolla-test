package controller

import (
	"github.com/denisqq/xsolla-test/app/service"
	"github.com/denisqq/xsolla-test/app/web/mapper"
	"github.com/denisqq/xsolla-test/app/web/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

//TODO В задание не сказано, что статистика должна быть по пользовательским ссылкам
type StatController struct {
	service.StatService
}

func NewStatController(statService service.StatService) StatController {
	return StatController{StatService: statService}
}

func (controller *StatController) Top(c *gin.Context) {
	var req request.TopRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	links, _ := controller.StatService.GetTop(req.Limit)
	c.JSON(http.StatusOK, gin.H{"links": mapper.ToLinkDtoList(links)})
}

func (controller *StatController) Conversion(c *gin.Context) {
	var req request.StatConversionGraphRequest
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	conversionGraph, err := controller.StatService.GetConversionGraph(req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusBadRequest, gin.H{"graph": conversionGraph})
}
