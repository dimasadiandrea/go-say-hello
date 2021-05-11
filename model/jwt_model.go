package model

type DecodedStructure struct {
	AdminId  uint32 `json:"admin_id"`
	ImageUrl string `json:"image_url"`
	Name     string `json:"name"`
	NIK      string `json:"nik"`
	Role     string `json:"role"`
	Username string `json:"username"`
}
