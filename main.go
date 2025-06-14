package main

import (
	"Vblog/test"
	"github.com/gin-gonic/gin"
	"github.com/infraboard/mcube/v2/ioc/config/http"
	"log"
)

func main() {
	test.DevelopmentSetup()
	server := gin.Default()
	if err := server.Run(http.Get().Addr()); err != nil {
		log.Println(err)
	}
}
