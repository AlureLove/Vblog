package main

import (
	blogApi "Vblog/apps/blog/api"
	blogImpl "Vblog/apps/blog/impl"
	"Vblog/apps/token"
	tokenApi "Vblog/apps/token/api"
	"Vblog/test"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/ioc/config/http"
	"log"
)

func main() {
	test.DevelopmentSetup()
	server := gin.Default()
	tokenApi.NewTokenApiHandler(token.GetService()).Registry(server)
	blogApi.NewBlogApiHandler(&blogImpl.BlogServiceImpl{}).Registry(server)
	if err := server.Run(http.Get().Addr()); err != nil {
		log.Println(err)
	}
}
