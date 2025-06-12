package blog

import (
	"Vblog/utils"
	"time"
)

type Blog struct {
	utils.ResourceMeta
	CreateBlogRequest
	Status
}

func (b *Blog) TableName() string {
	return "blogs"
}

type CreateBlogRequest struct {
	Title    string            `json:"title" gorm:"column:title;type:varchar(200)"`
	Summary  string            `json:"summary" gorm:"column:summary;type:text"`
	Content  string            `json:"content" gorm:"column:content;type:text"`
	Category string            `json:"category" gorm:"column:category;type:varchar(200);index"`
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
	Total int32   `json:"total"`
	Items []*Blog `json:"items"`
}
