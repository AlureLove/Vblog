package token

import (
	"github.com/google/uuid"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"time"
)

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

func NewToken(refUserId string) *Token {
	aExpiredAt := time.Now().AddDate(0, 0, 1)
	rExpiredAt := time.Now().AddDate(0, 0, 7)
	return &Token{
		RefUserId:            refUserId,
		AccessToken:          uuid.NewString(),
		AccessTokenExpireAt:  &aExpiredAt,
		IssuedAt:             time.Now(),
		RefreshToken:         uuid.NewString(),
		RefreshTokenExpireAt: &rExpiredAt,
	}
}

func (t *Token) SetRefUserName(refUserName string) *Token {
	t.RefUserName = refUserName
	return t
}

func (t *Token) String() string {
	return pretty.ToJSON(t)
}

func (t *Token) IsAccessTokenExpired() error {
	if t.AccessTokenExpireAt == nil {
		return nil
	}

	if time.Now().After(*t.AccessTokenExpireAt) {
		return exception.NewAccessTokenExpired("%s access token expired", t.AccessToken)
	}

	return nil
}

func (t *Token) IsRefreshTokenExpired() error {
	if t.RefreshTokenExpireAt == nil {
		return nil
	}

	if time.Now().After(*t.RefreshTokenExpireAt) {
		return exception.NewAccessTokenExpired("%s access token expired", t.RefreshToken)
	}

	return nil
}

func (t *Token) AccessTokenExpireTTL() int {
	if t.AccessTokenExpireAt == nil {
		return 0
	}
	return int(time.Until(*t.AccessTokenExpireAt).Seconds())
}

func (t *Token) TableName() string {
	return "tokens"
}
