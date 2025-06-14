package token

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/v2/ioc"
)

const (
	AppName = "token"
)

var (
	v = validator.New()
)

type Service interface {
	UserService
	InnerService
}

type UserService interface {
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	RevokeToken(context.Context, *RevokeTokenRequest) (*Token, error)
}

type InnerService interface {
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
}

type IssueTokenRequest struct {
	Username   string `json:"username" validate:"required"`
	Password   string `json:"password" validate:"required"`
	RememberMe bool   `json:"remember_me"`
}

type RevokeTokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type ValidateTokenRequest struct {
	AccessToken string `json:"access_token"`
}

func NewIssueTokenRequest(username, password string) *IssueTokenRequest {
	return &IssueTokenRequest{
		Username: username,
		Password: password,
	}
}

func NewValidateTokenRequest(accessToken string) *ValidateTokenRequest {
	return &ValidateTokenRequest{
		AccessToken: accessToken,
	}
}

func GetService() Service {
	return ioc.Controller().Get(AppName).(Service)
}

func (i *IssueTokenRequest) Validate() error {
	return v.Struct(i)
}
