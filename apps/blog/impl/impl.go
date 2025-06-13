package impl

import (
	"Vblog/apps/blog"
	"Vblog/middleware"
	"context"
	"fmt"
	"github.com/infraboard/mcube/v2/ioc/config/datasource"
)

var BlogService blog.Service = &BlogServiceImpl{}

type BlogServiceImpl struct {
}

func (b *BlogServiceImpl) CreateBlog(ctx context.Context, req *blog.CreateBlogRequest) (*blog.Blog, error) {
	ins, err := blog.NewBlog(req)
	if err != nil {
		return nil, err
	}

	tk := middleware.GetTokenFromCtx(ctx)
	ins.CreatedBy = tk.RefUserName
	
	if err = datasource.DBFromCtx(ctx).Create(ins).Error; err != nil {
		return nil, err
	}

	return ins, nil
}

func (b *BlogServiceImpl) QueryBlog(ctx context.Context, req *blog.QueryBlogRequest) (*blog.BlogSet, error) {
	query := datasource.DBFromCtx(ctx).Model(&blog.Blog{})
	if req.Keywords != "" {
		query = query.Where("title LIKE ?", "%"+req.Keywords+"%")
	}
	if req.Stage != nil {
		query = query.Where("stage = ?", req.Stage)
	}
	if req.CreatedBy != "" {
		query = query.Where("created_by = ?", req.CreatedBy)
	}
	if req.Category != "" {
		query = query.Where("category = ?", req.Category)
	}
	for k, v := range req.Tags {
		query = query.Where(fmt.Sprintf("tags->>'$.%s' = ?", k), v)
	}
	set := blog.NewBlogSet()

	if err := query.Count(&set.Total).Error; err != nil {
		return nil, err
	}

	err := query.Order("created_at DESC").Offset(int(req.Offset())).Limit(int(req.PageSize)).Find(&set.Items).Error
	if err != nil {
		return nil, err
	}

	return set, nil
}

func (b *BlogServiceImpl) DescribeBlog(ctx context.Context, req *blog.DescribeBlogRequest) (*blog.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BlogServiceImpl) UpdateBlog(ctx context.Context, req *blog.UpdateBlogRequest) (*blog.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BlogServiceImpl) PublishBlog(ctx context.Context, req *blog.PublishBlogRequest) (*blog.Blog, error) {
	//TODO implement me
	panic("implement me")
}

func (b *BlogServiceImpl) DeleteBlog(ctx context.Context, req *blog.DeleteBlogRequest) error {
	//TODO implement me
	panic("implement me")
}
