package user

import (
	"Vblog/utils"
	"time"
)

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

type Status struct {
	BlockReason string     `json:"block_reason" gorm:"column:block_reason;type:text"`
	BlockAt     *time.Time `json:"block_at" gorm:"column:block_at"`
}

func (s *Status) IsBlocked() bool {
	return s.BlockAt != nil
}

func (u *User) TableName() string {
	return "users"
}
