package api

import (
	"Vblog/apps/blog"
	"Vblog/middleware"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/http/gin/response"
	"github.com/infraboard/mcube/v2/ioc"
	iocgin "github.com/infraboard/mcube/v2/ioc/config/gin"
)

type BlogApiHandler struct {
	ioc.ObjectImpl
	blog blog.Service
}

func init() {
	ioc.Api().Registry(&BlogApiHandler{})
}

func (b *BlogApiHandler) Init() error {
	b.blog = blog.GetService()
	iocgin.RootRouter()

	r := iocgin.ObjectRouter(b)
	r.Use(middleware.Auth)
	r.POST("", b.CreateBlog)
	r.GET("", b.QueryBlog)
	return nil
}

func (b *BlogApiHandler) Name() string {
	return "blogs"
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
