package blog

import (
	"Vblog/utils"
	"context"
)

type Service interface {
	CreateBlog(context.Context, *CreateBlogRequest) (*Blog, error)
	QueryBlog(context.Context, *QueryBlogRequest) (*BlogSet, error)
	DescribeBlog(context.Context, *DescribeBlogRequest) (*Blog, error)
	UpdateBlog(context.Context, *UpdateBlogRequest) (*Blog, error)
	PublishBlog(context.Context, *PublishBlogRequest) (*Blog, error)
	DeleteBlog(context.Context, *DeleteBlogRequest) error
}

type QueryBlogRequest struct {
	utils.PageRequest
	Keywords string            `form:"keywords"`
	Stage    *STAGE            `json:"stage"`
	Username string            `json:"username"`
	Category string            `json:"category"`
	Tags     map[string]string `json:"tags"`
}

type DescribeBlogRequest struct {
	utils.GetRequest
}

type UpdateBlogRequest struct {
	utils.GetRequest
	CreateBlogRequest
}

type PublishBlogRequest struct {
	utils.GetRequest
	StatusSpec
}

type DeleteBlogRequest struct {
	utils.GetRequest
}
