package token

type Service interface {
	Outer
	Inner
}

type Outer interface {
	IssueToken()
	RevokeToken()
}

type Inner interface {
	ValidateToken()
}
