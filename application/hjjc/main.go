package main

import (
	"github.com/astaxie/beego"
	//"github.com/astaxie/beego/plugins/cors"

	"github.com/shiruitao/go-one/application/hjjc/initorm"
	_ "github.com/shiruitao/go-one/application/hjjc/routers"
	"github.com/astaxie/beego/plugins/cors"
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
