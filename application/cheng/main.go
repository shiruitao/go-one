package main

import (
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/cheng/initorm"
	_ "github.com/shiruitao/go-one/application/cheng/routers"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	initorm.InitMysql()
	beego.Run()
}
