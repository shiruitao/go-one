package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/cheng/models"
	"encoding/json"
	"github.com/shiruitao/go-one/application/cheng/log"
	"github.com/astaxie/beego/orm"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Create() {
	var admin models.Admin
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &admin)
	if err != nil {
		log.Logger.Error("json.Unmarshal", err)
	} else {
		err := models.AdminService.Create(admin)
		if err != nil {
			log.Logger.Error("models.Insert", err)
			this.Data["json"] = map[string]interface{}{"data": "false"}
		} else {
			this.Data["json"] = map[string]interface{}{"data": "true"}
		}
	}
	this.ServeJSON()
}

func (this *AdminController) Login() {
	var admin models.Admin
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &admin)
	if err != nil {
		log.Logger.Error("json.Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{"data": "false"}
	} else {
		flag, err := models.AdminService.Login(admin.Name, admin.Password)
		if err != nil {
			if err == orm.ErrNoRows{
				log.Logger.Debug("Invalid name!")
				this.Data["json"] = map[string]interface{}{"data": "name error"}
			} else {
				log.Logger.Error("出错", err)
				this.Data["json"] = map[string]interface{}{"data": "错误"}
			}
		} else {
			if !flag {
				log.Logger.Debug("Wrong Password!")
				this.Data["json"] = map[string]interface{}{"data":"密码出错"}
			} else {
				//this.SetSession("Name", admin.Name)
				this.Data["json"] = map[string]interface{}{"data": "成功"}
			}
		}
	}
	this.ServeJSON()
}