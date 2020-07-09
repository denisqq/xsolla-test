package bootstrap

import (
	"github.com/denisqq/xsolla-test/app/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

func connectDb() *gorm.DB {

	dbUrl := os.Getenv("DB_URL")
	database, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&model.Link{}, &model.User{}, &model.LinkHistory{})

	return database
}
