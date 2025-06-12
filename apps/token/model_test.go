package token_test

import (
	"Vblog/apps/token"
	"Vblog/test"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"testing"
)

func TestMigerate(t *testing.T) {
	if err := datasource.DB().AutoMigrate(&token.Token{}); err != nil {
		t.Fatal(err)
	}
}

func init() {
	test.DevelopmentSetup()
}
