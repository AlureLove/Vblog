package blog

import (
	"Vblog/utils"
	"context"
	"strings"
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
	Keywords  string            `json:"keywords" form:"keywords"`
	Stage     *STAGE            `json:"stage" form:"stage"`
	CreatedBy string            `json:"created_by" form:"created_by"`
	Category  string            `json:"category" form:"category"`
	Tags      map[string]string `json:"tags" form:"-"`
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

func NewQueryBlogRequest() *QueryBlogRequest {
	return &QueryBlogRequest{
		PageRequest: *utils.NewPageRequest(),
		Tags:        map[string]string{},
	}
}

func (q *QueryBlogRequest) SetTag(tag string) {
	kvItem := strings.Split(tag, ",")
	for i := range kvItem {
		kvString := kvItem[i]
		kv := strings.Split(kvString, "=")
		if len(kv) > 1 {
			q.Tags[kv[0]] = strings.Join(kv[1:], "=")
		}
	}
}
