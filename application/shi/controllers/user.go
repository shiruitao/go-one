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
	var user models.User
	//u := models.User{}
	err := this.ParseForm(&user)
	if err != nil {
		logger.Logger.Error("ParseForm:", err)
	}
	err = models.UserServer.Create(user)
	if err != nil {
		logger.Logger.Error("models.UserServer.Create:", err)
		this.Data["json"] = map[string]interface{}{"status": "success"}
	}

}
