package entity

import "gorm.io/gorm"

type Role struct {
	gorm.Model
	ID    uint64  `gorm:"primary_key:auto_increment"    json:"id"`
	Name  string  `gorm:"uniqueIndex;type:varchar(255)" json:"name"`
	Users *[]User `json:"user"`
}
