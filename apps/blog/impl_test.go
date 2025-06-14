package blog_test

import (
	"Vblog/apps/blog"
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
	ins, err := blog.GetService().CreateBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}

func TestQueryBlog(t *testing.T) {
	req := blog.NewQueryBlogRequest()
	ins, err := blog.GetService().QueryBlog(ctx, req)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(ins)
}
