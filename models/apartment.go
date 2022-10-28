package models

import "gorm.io/gorm"

type Apartment struct {
	gorm.Model
	ApartmentBacic
}

type ApartmentBacic struct {
	Father_id   int    `gorm:"column:father_id;type:int;" json:"father_id" form:"father_id"`
	Apt_id      string `gorm:"column:apt_id;type:varchar(15);" json:"apt_id" form:"apt_id"`
	Apt_name    string `gorm:"column:apt_name;type:varchar(63);" json:"apt_name" form:"apt_name"`
	Apt_address string `gorm:"column:apt_address;type:varchar(255);" json:"apt_address" form:"apt_address"`
	Apt_tel     string `gorm:"column:apt_tel;type:varchar(255);" json:"apt_tel" form:"apt_tel"`
}

func (table Apartment) TableName() string {
	return "apartment"
}

func GetAptList(keyword string) *gorm.DB {
	return DB.Model(new(Apartment)).
		Where("apt_name like ? OR apt_id like ?", "%"+keyword+"%", "%"+keyword+"%")
}
