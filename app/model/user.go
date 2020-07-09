package model

type User struct {
	BaseModel
	Username string `gorm:"column:username;not null"`
	Password string `gorm:"column:password;not null"`
}
