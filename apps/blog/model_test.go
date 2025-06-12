package blog_test

import (
	"Vblog/apps/blog"
	"Vblog/test"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
	"testing"
)

func TestMigerate(t *testing.T) {
	if err := datasource.DB().AutoMigrate(&blog.Blog{}); err != nil {
		t.Fatal(err)
	}
}

func init() {
	test.DevelopmentSetup()
}
