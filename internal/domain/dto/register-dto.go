package dto

//RegisterDTO is used when client post from /register url
type RegisterDTO struct {
	Name     string    `json:"name"             form:"name"      binding:"required"`
	Email    string    `json:"email"            form:"email"     binding:"required,email" `
	Password string    `json:"password"         form:"password"  binding:"required"`
	Role     string    `json:"role_name"        form:"role_name" binding:"required"`
	Driver   DriverDTO `json:"driver,omitempty" form:"driver"    binding:"required"`
}

type DriverDTO struct {
	DriverFile  uint64 `json:"driver_file"   form:"driver_file"   binding:"required"`
	Description string `json:"description"   form:"description"`
}
