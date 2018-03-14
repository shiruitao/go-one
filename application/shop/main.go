package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	//"github.com/astaxie/beego/plugins/cors"

	"github.com/shiruitao/go-one/application/shop/models"
	"github.com/shiruitao/go-one/application/shop/initorm"
	_ "github.com/shiruitao/go-one/application/shop/routers"
)

func Table() {
	force := false
	verbose := true
	_ = orm.RunSyncdb("default", force, verbose)
}

func main() {
	//if beego.BConfig.RunMode == "dev" {
	//	beego.BConfig.WebConfig.DirectoryIndex = true
	//	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	//}

	//beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	//	AllowOrigins:     []string{"*"},
	//	AllowMethods:     []string{"POST", "GET"},
	//	AllowHeaders:     []string{"Origin", "Content-Type"},
	//	ExposeHeaders:    []string{"Content-Length"},
	//	AllowCredentials: true,
	//}))

	//检查用户是否登录
	//beego.InsertFilter("/*", beego.BeforeRouter, filters.LoginFilter)
	orm.RegisterModel(new(models.User))
	initorm.InitMysql()
	Table()
	beego.Run()
}

