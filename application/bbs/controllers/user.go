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
	var user models.UserInfo

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		_, err := models.UserService.CreateUser(&user)
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
		user, _, err := models.UserService.Login(user.Name, user.Password)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.SetSession(common.SessionUserID, user.ID)
			this.SetSession(common.SessionUserName, user.Name)
			this.SetSession(common.SessionIsAdmin, user.IsAdmin)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "username": user.Name, common.SessionIsAdmin: user.IsAdmin, "avatar": user.Avatar}
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
		name := this.GetSession(common.SessionUserName)
		_, flag, err := models.UserService.Login(name.(string), user.Oldpass)
		if err != nil {
			log.Logger.Error("Old Password:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			if !flag {
				log.Logger.Debug("Wrong Password!")
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrWrongPass}
			} else {
				id := this.GetSession(common.SessionUserID)
				err := models.UserService.ChangePassword(id.(uint32), user.Newpass)
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

func (this *UserController) Get() {
	user, err := models.UserService.Get()
	if err != nil {
		log.Logger.Error("Errsql:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": user}
	}
	this.ServeJSON()
}

func (this *UserController) Delete() {
	var id struct{
		ID uint32 `json:"id"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Logger.Error("json.Unmarshal:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		isAdmin := this.GetSession(common.SessionIsAdmin).(bool)
		if isAdmin {
			_, err := models.UserService.Delete(id.ID)
			if err != nil {
				log.Logger.Error("Errsql:", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			} else {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrPermission}
		}
	}
	this.ServeJSON()
}
