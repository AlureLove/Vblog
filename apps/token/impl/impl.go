package impl

import (
	"Vblog/apps/token"
	"context"
)

var TokenService token.Service = &TokenServiceImpl{}

type TokenServiceImpl struct {
}

func (t *TokenServiceImpl) IssueToken(ctx context.Context, token *token.IssueTokenRequest) (*token.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TokenServiceImpl) RevokeToken(ctx context.Context, token *token.RevokeTokenRequest) (*token.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TokenServiceImpl) ValidateToken(ctx context.Context, token *token.ValidateTokenRequest) (*token.Token, error) {
	//TODO implement me
	panic("implement me")
}
