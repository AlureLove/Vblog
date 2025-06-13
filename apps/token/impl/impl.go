package impl

import (
	"Vblog/apps/token"
	"Vblog/apps/user"
	"Vblog/apps/user/impl"
	"context"
	"fmt"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
)

var TokenService token.Service = &TokenServiceImpl{
	user: impl.UserService,
}

type TokenServiceImpl struct {
	user user.AdminService
}

func (t *TokenServiceImpl) IssueToken(ctx context.Context, req *token.IssueTokenRequest) (*token.Token, error) {
	if err := req.Validate(); err != nil {
		return nil, exception.NewBadRequest("validate token error: %s", err)
	}

	u, err := t.user.DescribeUser(ctx, &user.DescribeUserRequest{
		DescribeBy: user.DESCRIBE_BY_USERNAME,
		Value:      req.Username,
	})
	if err != nil {
		return nil, err
	}

	if err = u.CheckPassword(req.Password); err != nil {
		return nil, err
	}

	tk := token.NewToken(fmt.Sprintf("%d", u.Id)).SetRefUserName(u.Username)

	if err = datasource.DBFromCtx(ctx).Create(tk).Error; err != nil {
		return nil, err
	}

	return tk, nil
}

func (t *TokenServiceImpl) RevokeToken(ctx context.Context, req *token.RevokeTokenRequest) (*token.Token, error) {
	//TODO implement me
	panic("implement me")
}

func (t *TokenServiceImpl) ValidateToken(ctx context.Context, req *token.ValidateTokenRequest) (*token.Token, error) {
	tk := &token.Token{}
	if err := datasource.DBFromCtx(ctx).Where("access_token = ?", req.AccessToken).Take(tk).Error; err != nil {
		return nil, err
	}

	if err := tk.IsAccessTokenExpired(); err != nil {
		return nil, err
	}
	
	return tk, nil
}
