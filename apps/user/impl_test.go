package user_test

import (
	"Vblog/apps/user"
	"Vblog/apps/user/impl"
	"context"
	"testing"
)

var (
	ctx = context.Background()
)

func TestRegistry(t *testing.T) {
	ins, err := impl.UserService.Registry(ctx, &user.RegistryRequest{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
