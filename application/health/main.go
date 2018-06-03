package main

import (
	_ "github.com/shiruitao/go-one/application/health/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"

	"github.com/shiruitao/go-one/application/health/initorm"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	initorm.InitMysql()
	beego.Run()
}
