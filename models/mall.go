package models

import (
	"time"

	"gorm.io/gorm"
)

type Mall struct {
	MallBasic
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type MallBasic struct {
	ID           uint   `gorm:"primarykey" json:"ID" form:"ID" example:"666"`
	Mall_id      string `gorm:"column:mall_id;type:varchar(15);" json:"mall_id" form:"mall_id" example:"dingzhen001"`
	Mall_name    string `gorm:"column:mall_name;type:varchar(63);" json:"mall_name" form:"mall_name" example:"丁真烟酒专卖超商"`
	Mall_address string `gorm:"column:mall_address;type:varchar(255);" json:"mall_address" form:"mall_address" example:"四川省甘孜藏族自治州理塘县朱雀大道1-1"`
	Mall_tel     string `gorm:"column:mall_tel;type:varchar(255);" json:"mall_tel" form:"mall_tel" example:"400-114514"`
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
