package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/shi/models"
	"github.com/shiruitao/go-one/libs/logger"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Create() {
	u := models.User{}
	if err := this.ParseForm(&u); err != nil {
		logger.Logger.Error("ParseForm:", err)
	}
	err := models.UserServer.Create(u)
	if err != nil {
		logger.Logger.Error("models.UserServer.Create:", err)
		this.Data["json"] = map[string]interface{}{"status": "success"}
	}

}
