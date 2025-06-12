package token

import "time"

type Token struct {
	Id                   uint       `json:"id" gorm:"primary_key;column:id"`
	RefUserId            string     `json:"ref_user_id" gorm:"column:ref_user_id"`
	AccessToken          string     `json:"access_token" gorm:"column:access_token;unique;index"`
	AccessTokenExpireAt  *time.Time `json:"access_token_expire_at" gorm:"column:access_token_expire_at"`
	IssuedAt             time.Time  `json:"issued_at" gorm:"column:issued_at"`
	RefreshToken         string     `json:"refresh_token" gorm:"column:refresh_token;unique;index"`
	RefreshTokenExpireAt *time.Time `json:"refresh_token_expire_at" gorm:"column:refresh_token_expire_at"`
	RefUserName          string     `json:"ref_user_name" gorm:"-"`
}

func (t *Token) TableName() string {
	return "tokens"
}
