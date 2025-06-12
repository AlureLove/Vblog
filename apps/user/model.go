package user

import "Vblog/utils"

type User struct {
	utils.ResourceMeta
	RegistryRequest
}

type RegistryRequest struct {
	Username string `json:"username" gorm:"column:username;unique;index"`
	Password string `json:"password" gorm:"column:password;type:varchar(255)"`
	Profile
}

type Profile struct {
	Avatar   string `json:"avatar" gorm:"column:avatar;type:varchar(255)"`
	Nickname string `json:"nickname" gorm:"column:nickname;type:varchar(100)"`
	Email    string `json:"email" gorm:"column:email;type:varchar(100)"`
}

func (u *User) TableName() string {
	return "users"
}
