package service

import (
	"github.com/denisqq/xsolla-test/app/dao"
	"github.com/denisqq/xsolla-test/app/model"
	"github.com/denisqq/xsolla-test/app/utils"
	"github.com/gofrs/uuid"
	"os"
)

type LinkService struct {
	LinkRepository dao.LinkRepository
}

func NewLinkService(linkRepository dao.LinkRepository) LinkService {
	return LinkService{LinkRepository: linkRepository}
}

func (service *LinkService) CreateShortLink(link model.Link) model.Link {
	shortUrl := utils.GenerateRandomString()
	_, err := service.LinkRepository.FindByShortUrl(shortUrl)
	if err != nil {
		link.ShortUrl = shortUrl
		link.Link = os.Getenv("REDIRECT_URL") + "/" + shortUrl
		return service.LinkRepository.Create(link)
	}
	return service.CreateShortLink(link)
}

func (service *LinkService) GetShortLinkById(id uuid.UUID) (model.Link, error) {
	return service.LinkRepository.FindByLinkId(id)
}

func (service *LinkService) GetUserLinks(username string) ([]model.Link, error) {
	return service.LinkRepository.FindLinksByUser(username)
}

func (service *LinkService) GetShortLink(shortLink string) (model.Link, error) {
	return service.LinkRepository.FindByShortUrl(shortLink)
}

func (service *LinkService) IncreaseLinkConversation(link model.Link) {
	_ = service.LinkRepository.IncreaseLinkConversation(link)
}

func (service *LinkService) DeleteLink(linkId uuid.UUID) error {
	return service.LinkRepository.DeleteLink(linkId)
}