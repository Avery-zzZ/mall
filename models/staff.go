package models

import "gorm.io/gorm"

type Staff struct {
	gorm.Model
	Father_id  int     `gorm:"column:father_id;type:int;" json:"father_id"`
	Staff_id   string  `gorm:"column:staff_id;type:varchar(15);" json:"staff_id"`
	Staff_name string  `gorm:"column:staff_name;type:varchar(63);" json:"staff_name"`
	Staff_pos  string  `gorm:"column:staff_pos;type:varchar(255);" json:"staff_address"`
	Staff_tel  string  `gorm:"column:staff_tel;type:varchar(255);" json:"staff_tel"`
	Staff_sal  float64 `gorm:"column:staff_sal;type:decimal(10,2);" json:"staff_sal"`
}

type StaffBasic struct {
	ID         uint    `gorm:"primarykey"`
	Father_id  int     `gorm:"column:father_id;type:int;" json:"father_id"`
	Staff_id   string  `gorm:"column:staff_id;type:varchar(15);" json:"staff_id"`
	Staff_name string  `gorm:"column:staff_name;type:varchar(63);" json:"staff_name"`
	Staff_pos  string  `gorm:"column:staff_pos;type:varchar(255);" json:"staff_address"`
	Staff_tel  string  `gorm:"column:staff_tel;type:varchar(255);" json:"staff_tel"`
	Staff_sal  float64 `gorm:"column:staff_sal;type:decimal(10,2);" json:"staff_sal"`
}

func (table Staff) TableName() string {
	return "staff"
}

func GetStaffList(keyword string) *gorm.DB {
	return DB.Model(new(Staff)).
		Where("staff_name like ? OR staff_id like ?", "%"+keyword+"%", "%"+keyword+"%")
}
