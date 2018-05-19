package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/bbs/common"
	"github.com/shiruitao/go-one/application/bbs/log"
	"github.com/shiruitao/go-one/application/bbs/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) CreateUser() {
	var user struct {
		Name            string `json:"name"`
		Nickname        string `json:"nickname"`
		Password        string `json:"password"`
		Avatar			string `json:"avatar"`
		ConfirmPassword string `json:"confirm_password"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		_, err := models.UserService.CreateUser(user.Name, user.Nickname, user.Password, user.Avatar, user.ConfirmPassword)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}
func (this *UserController) Login() {

	var user struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		isAdmin, _, err := models.UserService.Login(user.Name, user.Password)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.SetSession(common.SessionUser, user.Name)
			this.SetSession(common.SessionIsAdmin, isAdmin)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, common.SessionUser: user.Name, common.SessionIsAdmin: isAdmin}
		}
	}
	this.ServeJSON()
}

func (this *UserController) ChangePassword() {
	var user struct {
		Oldpass string
		Newpass string
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		log.Logger.Error("json.Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		name := this.GetSession(common.SessionUser)
		_, flag, err := models.UserService.Login(name.(string), user.Oldpass)
		if err != nil {
			log.Logger.Error("Old Password:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			if !flag {
				log.Logger.Debug("Wrong Password!")
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrWrongPass}
			} else {
				err := models.UserService.ChangePassword(name.(string), user.Newpass)
				if err != nil {
					log.Logger.Error("models.ChangePass:", err)
					this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
				} else {
					this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
				}
			}
		}
	}
	this.ServeJSON()
}
