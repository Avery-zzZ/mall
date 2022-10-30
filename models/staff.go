package models

import (
	"time"

	"gorm.io/gorm"
)

type Staff struct {
	StaffBasic
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type StaffBasic struct {
	ID         uint    `gorm:"primarykey" json:"ID" form:"ID" example:"114514"`
	Father_id  int     `gorm:"column:father_id;type:int;" json:"father_id" form:"father_id" example:"8848"`
	Staff_id   string  `gorm:"column:staff_id;type:varchar(15);" json:"staff_id" form:"staff_id" example:"dz1r6-zhenzhu"`
	Staff_name string  `gorm:"column:staff_name;type:varchar(63);" json:"staff_name" form:"staff_name" example:"珍珠"`
	Staff_pos  string  `gorm:"column:staff_pos;type:varchar(255);" json:"staff_address" form:"staff_address" example:"销售经理"`
	Staff_tel  string  `gorm:"column:staff_tel;type:varchar(255);" json:"staff_tel" form:"staff_tel" example:"18355555555"`
	Staff_sal  float64 `gorm:"column:staff_sal;type:decimal(10,2);" json:"staff_sal" form:"staff_sal" example:"100000"`
}

func (table Staff) TableName() string {
	return "staff"
}

func GetStaffList(keyword string) *gorm.DB {
	return DB.Model(new(Staff)).
		Where("staff_name like ? OR staff_id like ?", "%"+keyword+"%", "%"+keyword+"%")
}
