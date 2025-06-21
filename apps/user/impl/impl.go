package impl

import (
	"Vblog/apps/user"
	"context"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/ioc"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var _ user.Service = (*UserServiceImpl)(nil)

type UserServiceImpl struct {
	ioc.ObjectImpl
}

func init() {
	ioc.Controller().Registry(&UserServiceImpl{})
}

func (u *UserServiceImpl) Name() string {
	return user.AppName
}

func (u *UserServiceImpl) Registry(ctx context.Context, req *user.RegistryRequest) (*user.User, error) {
	ins, err := user.NewUser(req)
	if err != nil {
		return nil, err
	}

	hashPwd, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	ins.Password = string(hashPwd)

	if err = datasource.DBFromCtx(ctx).Create(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (u *UserServiceImpl) UpdatePassword(ctx context.Context, req *user.UpdatePasswordRequest) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) RestPassword(ctx context.Context, req *user.ResetPasswordRequest) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) UpdateProfile(ctx context.Context, req *user.UpdateProfileRequest) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) UnRegistry(ctx context.Context, req *user.UnRegistryRequest) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) UpdateUserStatus(ctx context.Context, req *user.UpdateUserStatusRequest) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) DescribeUser(ctx context.Context, req *user.DescribeUserRequest) (*user.User, error) {
	query := datasource.DBFromCtx(ctx)
	switch req.DescribeBy {
	case user.DESCRIBE_BY_ID:
		query = query.Where("id = ?", req.Value)
	case user.DESCRIBE_BY_USERNAME:
		query = query.Where("username = ?", req.Value)
	}

	ins := &user.User{}
	if err := query.Take(ins).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, exception.NewNotFound("user not found")
		}
		return nil, err
	}

	return ins, nil
}
