package controller

import (
	"github.com/denisqq/xsolla-test/app/model"
	"github.com/denisqq/xsolla-test/app/service"
	"github.com/denisqq/xsolla-test/app/web/mapper"
	"github.com/denisqq/xsolla-test/app/web/request"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"net/http"
)

type LinkController struct {
	LinkService service.LinkService
	UserService service.UserService
}

func NewLinkController(linkService service.LinkService, userService service.UserService) LinkController {
	return LinkController{LinkService: linkService, UserService: userService}
}

func (controller *LinkController) CreateShortLink(c *gin.Context) {
	username := c.GetString(gin.AuthUserKey)
	user, _ := controller.UserService.FindByUsername(username)

	var input request.CreateLinkRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		toError(c, err)
		return
	}

	link := controller.LinkService.CreateShortLink(model.Link{UserId: user.ID, OriginUrl: input.Url})
	c.JSON(http.StatusOK, gin.H{"link": mapper.ToLinkDto(link)})
}

func (controller *LinkController) GetLink(c *gin.Context) {
	linkId, err := uuid.FromString(c.Param("linkId"))
	if err != nil {
		toError(c, err)
		return
	}
	link, err := controller.LinkService.GetShortLinkById(linkId)
	if err != nil {
		toError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"link": mapper.ToLinkDto(link)})
}

func (controller *LinkController) GetSelfLinks(c *gin.Context) {
	username := c.GetString(gin.AuthUserKey)

	links, err := controller.LinkService.GetUserLinks(username)
	if err != nil {
		toError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"links": mapper.ToLinkDtoList(links)})
}

func (controller *LinkController) DeleteLink(c *gin.Context)  {
	linkId, err := uuid.FromString(c.Param("linkId"))
	if err != nil {
		toError(c, err)
		return
	}
	err = controller.LinkService.DeleteLink(linkId)
	if err != nil {
		toError(c, err)
		return
	}
	c.Writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	c.Writer.WriteHeader(http. StatusNoContent)
	//c.JSON(204, nil)
}

func (controller *LinkController) RedirectToUrl(c *gin.Context) {
	shortUrl := c.Param("shortUrl")
	link, err := controller.LinkService.GetShortLink(shortUrl)
	if err != nil {
		toError(c, err)
		return
	}
	c.Redirect(http.StatusMovedPermanently, link.OriginUrl)
}

func toError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
}
