package blog

import (
	"Vblog/utils"
	"github.com/go-playground/validator/v10"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/tools/pretty"
	"time"
)

var (
	v = validator.New()
)

type Blog struct {
	utils.ResourceMeta
	CreateBlogRequest
	Status
}

type CreateBlogRequest struct {
	Title    string            `json:"title" gorm:"column:title;type:varchar(200)" validate:"required"`
	Summary  string            `json:"summary" gorm:"column:summary;type:text" validate:"required"`
	Content  string            `json:"content" gorm:"column:content;type:text" validate:"required"`
	Category string            `json:"category" gorm:"column:category;type:varchar(200);index" validate:"required"`
	Tags     map[string]string `json:"tags" gorm:"column:tags;serializer:json"`
}

type Status struct {
	StatusSpec
	ChangedAt *time.Time `json:"changed_at" gorm:"column:changed_at"`
}

type StatusSpec struct {
	// 0: draft 1: published 2: checking...
	Stage STAGE `json:"stage" gorm:"column:stage;type:tinyint(1);index"`
}

type BlogSet struct {
	Total int64   `json:"total"`
	Items []*Blog `json:"items"`
}

func NewBlog(req *CreateBlogRequest) (*Blog, error) {
	if err := v.Struct(req); err != nil {
		return nil, exception.NewBadRequest("validate blog error: %s", err)
	}
	return &Blog{
		ResourceMeta:      *utils.NewResourceMeta(),
		CreateBlogRequest: *req,
	}, nil
}

func NewBlogSet() *BlogSet {
	return &BlogSet{
		Items: []*Blog{},
	}
}

func (c *CreateBlogRequest) Validate() error {
	return v.Struct(c)
}

func (b *Blog) String() string {
	return pretty.ToJSON(b)
}

func (b *Blog) TableName() string {
	return "blogs"
}
