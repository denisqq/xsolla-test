package mapper

import (
	"github.com/denisqq/xsolla-test/app/model"
	"github.com/denisqq/xsolla-test/app/web/dto"
)

func ToLinkDto(link model.Link) dto.LinkDto {
	return dto.LinkDto{ID: link.ID, Link: link.Link, Conversion: link.Conversion}
}

func ToLinkDtoList(links []model.Link) []dto.LinkDto {
	var linksDto []dto.LinkDto
	for _, link := range links {
		linksDto = append(linksDto, dto.LinkDto{ID: link.ID, Link: link.Link, Conversion: link.Conversion})
	}
	return linksDto
}
