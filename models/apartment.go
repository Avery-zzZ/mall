package models

import (
	"time"

	"gorm.io/gorm"
)

type Apartment struct {
	ApartmentBasic
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ApartmentBasic struct {
	ID          uint   `gorm:"primarykey" json:"ID" example:"8848"`
	Father_id   int    `gorm:"column:father_id;type:int;" json:"father_id" form:"father_id" example:"666"`
	Apt_id      string `gorm:"column:apt_id;type:varchar(15);" json:"apt_id" form:"apt_id" example:"dz1-relx06"`
	Apt_name    string `gorm:"column:apt_name;type:varchar(63);" json:"apt_name" form:"apt_name" example:"悦刻销售部"`
	Apt_address string `gorm:"column:apt_address;type:varchar(255);" json:"apt_address" form:"apt_address" example:"超商B区二期三楼北面电梯旁"`
	Apt_tel     string `gorm:"column:apt_tel;type:varchar(255);" json:"apt_tel" form:"apt_tel" example:"18355555555"`
}

func (table Apartment) TableName() string {
	return "apartment"
}

func GetAptList(keyword string) *gorm.DB {
	return DB.Model(new(Apartment)).
		Where("apt_name like ? OR apt_id like ?", "%"+keyword+"%", "%"+keyword+"%")
}
