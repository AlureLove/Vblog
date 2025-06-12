package impl

import (
	"Vblog/apps/blog"
	"context"
)

var BlogService blog.Service = &BlogServiceImpl{}

type BlogServiceImpl struct {
}

func (b *BlogServiceImpl) CreateBlog(ctx context.Context, blog *blog.CreateBlogRequest) (*blog.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BlogServiceImpl) QueryBlog(ctx context.Context, request *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BlogServiceImpl) DescribeBlog(ctx context.Context, request *blog.DescribeBlogRequest) (*blog.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BlogServiceImpl) UpdateBlog(ctx context.Context, request *blog.UpdateBlogRequest) (*blog.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BlogServiceImpl) PublishBlog(ctx context.Context, request *blog.PublishBlogRequest) (*blog.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BlogServiceImpl) DeleteBlog(ctx context.Context, request *blog.DeleteBlogRequest) error {
	//TODO implement me
	panic("implement me")
}
