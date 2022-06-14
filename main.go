package main

import (
	_ "beegotest/routers"
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// 判断beego.DEV是否为 'dev'，设置swagger的路径并启用
	if beego.DEV == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	fmt.Println(beego.AppConfig)

	beego.Run()
}
