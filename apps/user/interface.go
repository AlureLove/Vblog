package user

import "context"

type Service interface {
	UserService
	AdminService
}

type AdminService interface {
	UpdateUserStatus(context.Context, *UpdateUserStatusRequest) (*User, error)
	DescribeUser(context.Context, *DescribeUserRequest) (*User, error)
}

type UserService interface {
	Registry(context.Context, *RegistryRequest) (*User, error)
	UpdatePassword(context.Context, *UpdatePasswordRequest) error
	RestPassword(context.Context, *ResetPasswordRequest) error
	UpdateProfile(context.Context, *UpdateProfileRequest) (*User, error)
	UnRegistry(context.Context, *UnRegistryRequest)
}

type DescribeUserRequest struct {
	DescribeBy DESCRIBE_BY `json:"describe_by"`
	Value      string      `json:"value"`
}

type UpdateUserStatusRequest struct {
	UserId string `json:"user_id"`
	Status
}

type UpdatePasswordRequest struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ResetPasswordRequest struct {
	Username    string `json:"username"`
	NewPassword string `json:"new_password"`
	VerifyCode  string `json:"verify_code"`
}

type UpdateProfileRequest struct {
	UserId string `json:"user_id"`
	Profile
}

type UnRegistryRequest struct {
	Username string `json:"username"`
}
