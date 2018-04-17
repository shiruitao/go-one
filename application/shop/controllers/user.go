package controllers

import (
	"log"
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/shop/models"
	"github.com/shiruitao/go-one/application/shop/common"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) CreateUser() {
	var user models.User

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		log.Println(err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	}
	_, err = models.UserService.CreateUser(&user)
	if err != nil {
		log.Println(err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	}
	this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
	this.ServeJSON()
}
