package entity

import (
	"gorm.io/gorm"
	"time"
)

type Trip struct {
	gorm.Model
	ID        uint64    `gorm:"primary_key:auto_increment"                                       json:"id"`
	StartTime time.Time `gorm:"type:time; not null"                                              json:"start_time"`
	EndTime   time.Time `gorm:"type:time"                                                        json:"end_time,omitempty"`
	DriverID  uint64    `gorm:""                                                                 json:"driver_id"`
	Finished  bool      `gorm:"type:bool;default:false"                                          json:"finished"`
	Driver    *Driver   `gorm:"foreignkey:DriverID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"driver"`
}
