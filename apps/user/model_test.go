package user_test

import (
	"Vblog/apps/user"
	"Vblog/test"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"testing"
)

func TestMigerate(t *testing.T) {
	if err := datasource.DB().AutoMigrate(&user.User{}); err != nil {
		t.Fatal(err)
	}
}

func init() {
	test.DevelopmentSetup()
}
