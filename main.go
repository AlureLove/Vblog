package main

import (
	_ "Vblog/apps"
	"github.com/infraboard/mcube/v2/ioc/server/cmd"
)

func main() {
	//test.DevelopmentSetup()
	//if err := server.Run(context.Background()); err != nil {
	//	panic(err)
	//}
	cmd.Start()
}
