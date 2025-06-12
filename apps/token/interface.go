package token

type Service interface {
	UserService
	InnerService
}

type UserService interface {
	IssueToken()
	RevokeToken()
}

type InnerService interface {
	ValidateToken()
}
