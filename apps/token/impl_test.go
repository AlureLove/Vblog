package token_test

import (
	"Vblog/apps/token"
	"Vblog/apps/token/impl"
	"context"
	"testing"
)

var (
	ctx = context.Background()
)

func TestIssueToken(t *testing.T) {
	req := token.NewIssueTokenRequest("admin", "123456")
	ins, err := impl.TokenService.IssueToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestValidateToken(t *testing.T) {
	req := token.NewValidateTokenRequest("1b70e2d4-af9a-4f74-8839-cf45ec696a54")
	ins, err := impl.TokenService.ValidateToken(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
