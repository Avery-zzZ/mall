package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"column:name;type:varchar(31);" json:"staff_id"`
	Password string `gorm:"column:passwd;type:varchar(127);" json:"staff_name"`
	Grade int `gorm:"column:grade;type:int;" json:"grade"`
}

func (table User) TableName() string {
	return "user"
}