package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/shi/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) Create() {
	models.UserServer.Create()
}
