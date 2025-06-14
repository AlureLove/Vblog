package api

import (
	"Vblog/apps/token"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/http/gin/response"
	"github.com/infraboard/mcube/v2/ioc"
	iocgin "github.com/infraboard/mcube/v2/ioc/config/gin"
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
	iocgin.RootRouter()
	r := iocgin.ObjectRouter(t)
	r.POST("", t.IssueToken)
	r.DELETE("", t.RevokeToken)
	return nil
}

func (t *TokenApiHandler) Name() string {
	return "tokens"
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
