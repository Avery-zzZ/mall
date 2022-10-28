package models

import (
	"gorm.io/gorm"
)

type Mall struct {
	gorm.Model
	MallBasic
}

type MallBasic struct {
	Mall_id      string `gorm:"column:mall_id;type:varchar(15);" json:"mall_id" form:"mall_id"`
	Mall_name    string `gorm:"column:mall_name;type:varchar(63);" json:"mall_name" form:"mall_name"`
	Mall_address string `gorm:"column:mall_address;type:varchar(255);" json:"mall_address" form:"mall_address"`
	Mall_tel     string `gorm:"column:mall_tel;type:varchar(255);" json:"mall_tel" form:"mall_tel"`
}

func (table Mall) TableName() string {
	return "mall"
}
func (table MallBasic) TableName() string {
	return "mall"
}

func GetMallList(keyword string) *gorm.DB {
	return DB.Model(new(Mall)).
		Where("mall_name like ? OR mall_id like ?", "%"+keyword+"%", "%"+keyword+"%")
}
