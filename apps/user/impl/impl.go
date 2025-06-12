package impl

import (
	"Vblog/apps/user"
	"context"
)

var UserService user.Service = &UserServiceImpl{}

type UserServiceImpl struct {
}

func (u *UserServiceImpl) Registry(ctx context.Context, user *user.RegistryRequest) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) UpdatePassword(ctx context.Context, user *user.UpdatePasswordRequest) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) RestPassword(ctx context.Context, user *user.ResetPasswordRequest) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) UpdateProfile(ctx context.Context, user *user.UpdateProfileRequest) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) UnRegistry(ctx context.Context, user *user.UnRegistryRequest) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) UpdateUserStatus(ctx context.Context, user *user.UpdateUserStatusRequest) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u *UserServiceImpl) DescribeUser(ctx context.Context, user *user.DescribeUserRequest) (*user.User, error) {
	//TODO implement me
	panic("implement me")
}
