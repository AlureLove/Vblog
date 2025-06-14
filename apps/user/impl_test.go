package user_test

import (
	"Vblog/apps/user"
	"context"
	"testing"
)

var (
	ctx = context.Background()
)

func TestRegistry(t *testing.T) {
	req := user.NewRegistryRequest()
	req.Username = "admin"
	req.Password = "123456"
	ins, err := user.GetService().Registry(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestDescribeUser(t *testing.T) {
	ins, err := user.GetService().DescribeUser(ctx, &user.DescribeUserRequest{
		user.DESCRIBE_BY_USERNAME, "admin",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
