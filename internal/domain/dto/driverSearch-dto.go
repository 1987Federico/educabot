package dto

type DriverSearch struct {
	Offset int `json:"offset" form:"offset"`
	Limit  int `json:"limit"  form:"limit"`
}
