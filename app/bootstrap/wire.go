//+build wireinject

package bootstrap

import (
	"github.com/denisqq/xsolla-test/app/dao"
	"github.com/denisqq/xsolla-test/app/security"
	"github.com/denisqq/xsolla-test/app/service"
	"github.com/denisqq/xsolla-test/app/web/controller"
	"github.com/google/wire"
	"github.com/jinzhu/gorm"
)

func InitUserController(db *gorm.DB) controller.UserController {
	wire.Build(dao.NewUserRepository, service.NewUserService, controller.NewUserController)
	return controller.UserController{}
}

func InitBasicAuth(db *gorm.DB) security.BasicAuth {
	wire.Build(dao.NewUserRepository, security.NewBasicAuthUserDetailsService, security.NewAuthenticate, security.NewBasicAuth)
	return security.BasicAuth{}
}

func InitLinkController(db *gorm.DB) controller.LinkController {
	wire.Build(dao.NewLinkRepository, dao.NewUserRepository, service.NewUserService, service.NewLinkService, controller.NewLinkController)
	return controller.LinkController{}
}

func InitLinkHistoryService(db *gorm.DB) service.LinkHistoryService {
	wire.Build(dao.NewLinkRepository, dao.NewLinkHistoryRepository, service.NewLinkService, service.NewLinkHistoryService)
	return service.LinkHistoryService{}
}

func InitStatController(db *gorm.DB) controller.StatController {
	wire.Build(dao.NewStatRepository, service.NewStatService, controller.NewStatController)
	return controller.StatController{}
}