package entity

//User represents users table in database
type User struct {
	ID       uint64  `gorm:"primary_key:auto_increment"                                     json:"id"`
	Name     string  `gorm:"type:varchar(255)"                                              json:"name"`
	Email    string  `gorm:"uniqueIndex;type:varchar(255)"                                  json:"email"`
	Password string  `gorm:"->;<-;not null"                                                 json:"-"`
	Token    string  `gorm:"-"                                                              json:"token,omitempty"`
	RoleID   uint64  `gorm:"not null"                                                       json:"role_id"`
	Roles    *Role   `gorm:"foreignkey:RoleID;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"role,omitempty"`
	Driver   *Driver `json:"driver,omitempty"`
}
