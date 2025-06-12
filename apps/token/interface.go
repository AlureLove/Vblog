package token

import "context"

type Service interface {
	UserService
	InnerService
}

type UserService interface {
	IssueToken(context.Context, *IssueTokenRequest) (*Token, error)
	RevokeToken(context.Context, *RevokeTokenRequest) (*Token, error)
}

type IssueTokenRequest struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type RevokeTokenRequest struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type InnerService interface {
	ValidateToken(context.Context, *ValidateTokenRequest) (*Token, error)
}

type ValidateTokenRequest struct {
	AccessToken string `json:"access_token"`
}
