package entity

//Driver struct represents books table in database
type Driver struct {
	ID          uint64 `gorm:"primary_key:auto_increment"                                     json:"id"`
	DriverFile  uint64 `gorm:"uniqueIndex"                                                    json:"driver_file"`
	Description string `gorm:"type:text"                                                      json:"description"`
	UserID      uint64 `gorm:"not null"                                                       json:"-"`
	User        User   `gorm:"foreignkey:UserID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}
