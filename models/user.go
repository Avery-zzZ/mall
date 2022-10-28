package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(31);" json:"username" form:"username"`
	Password string `gorm:"column:passwd;type:varchar(127);" json:"passwd" form:"passwd"`
	Grade int `gorm:"column:grade;type:int;" json:"grade" form:"grade"`
}

func (table User) TableName() string {
	return "user"
}