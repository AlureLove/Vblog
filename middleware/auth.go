package middleware

import (
	"Vblog/apps/token"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/exception"
	"github.com/infraboard/mcube/v2/http/gin/response"
	"strings"
)

type TokenCtxKey struct {
}

func Auth(g *gin.Context) {
	tk := g.GetHeader("Authorization")
	tkList := strings.Split(tk, " ")
	accessToken := ""
	if len(tkList) == 2 {
		accessToken = tkList[1]
	}

	t, err := token.GetService().ValidateToken(g.Request.Context(), token.NewValidateTokenRequest(accessToken))
	if err != nil {
		response.Failed(g, exception.NewUnauthorized("token check failed: %s", err))
		g.Abort()
		return
	}

	ctx := context.WithValue(g.Request.Context(), TokenCtxKey{}, t)
	g.Request = g.Request.WithContext(ctx)
}

func GetTokenFromCtx(ctx context.Context) *token.Token {
	return ctx.Value(TokenCtxKey{}).(*token.Token)
}
