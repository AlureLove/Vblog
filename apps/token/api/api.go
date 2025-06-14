package api

import (
	"Vblog/apps/token"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/http/gin/response"
	"github.com/infraboard/mcube/v2/ioc"
)

type TokenApiHandler struct {
	ioc.ObjectImpl
	token token.UserService
}

func init() {
	ioc.Api().Registry(&TokenApiHandler{})
}

func (t *TokenApiHandler) Init() error {
	t.token = token.GetService()
	return nil
}

func (t *TokenApiHandler) Registry(g *gin.Engine) {
	router := g.Group("/vblog/api/v1/tokens")
	router.POST("", t.IssueToken)
	router.DELETE("", t.RevokeToken)
}

func (t *TokenApiHandler) IssueToken(ctx *gin.Context) {
	req := &token.IssueTokenRequest{}
	if err := ctx.BindJSON(req); err != nil {
		response.Failed(ctx, err)
		return
	}

	ins, err := t.token.IssueToken(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}

func (t *TokenApiHandler) RevokeToken(ctx *gin.Context) {
	req := &token.RevokeTokenRequest{}
	if err := ctx.BindJSON(req); err != nil {
		response.Failed(ctx, err)
		return
	}

	ins, err := t.token.RevokeToken(ctx.Request.Context(), req)
	if err != nil {
		response.Failed(ctx, err)
		return
	}
	response.Success(ctx, ins)
}
