package api

import (
	"Vblog/apps/blog"
	"Vblog/middleware"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/http/gin/response"
)

type BlogApiHandler struct {
	blog blog.Service
}

func NewBlogApiHandler(blogImpl blog.Service) *BlogApiHandler {
	return &BlogApiHandler{
		blog: blogImpl,
	}
}

func (b *BlogApiHandler) Registry(g *gin.Engine) {
	router := g.Group("/vblog/api/v1/blogs")
	router.Use(middleware.Auth)
	router.POST("", b.CreateBlog)
	router.GET("", b.QueryBlog)
}

func (b *BlogApiHandler) CreateBlog(ctx *gin.Context) {
	req := &blog.CreateBlogRequest{}
	if err := ctx.BindJSON(req); err != nil {
		response.Failed(ctx, err)
		return
	}

	ins, err := b.blog.CreateBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}

func (b *BlogApiHandler) QueryBlog(ctx *gin.Context) {
	req := blog.NewQueryBlogRequest()
	if err := ctx.BindQuery(req); err != nil {
		response.Failed(ctx, err)
		return
	}
	req.SetTag(ctx.Query("tags"))

	ins, err := b.blog.QueryBlog(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}
