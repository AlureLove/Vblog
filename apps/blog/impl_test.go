package blog_test

import (
	"Vblog/apps/blog"
	"Vblog/apps/blog/impl"
	"context"
	"testing"
)

var (
	ctx = context.Background()
)

func TestCreateBlog(t *testing.T) {
	req := &blog.CreateBlogRequest{
		Title:    "Golang",
		Summary:  "Full-stack",
		Content:  "GORM + Gin",
		Category: "Programming",
	}
	ins, err := impl.BlogService.CreateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryBlog(t *testing.T) {
	req := blog.NewQueryBlogRequest()
	ins, err := impl.BlogService.QueryBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
