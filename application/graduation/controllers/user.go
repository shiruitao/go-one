package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/graduation/common"
	"github.com/shiruitao/go-one/application/graduation/log"
	"github.com/shiruitao/go-one/application/graduation/models"
)

type UserController struct {
	beego.Controller
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
			this.SetSession(common.SessionUserName, user.Name)
			this.SetSession(common.SessionUserNum, user.Number)
			this.SetSession(common.SessionUserID, user.ID)
			this.SetSession(common.SessionPower, user.Power)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "username": user.Name, common.SessionPower: user.Power}
		}
	}
	this.ServeJSON()
}

func (this *UserController) CreateUserInfo() {
	var (
		user  models.User
	)

	power := this.GetSession(common.SessionPower).(int8)
	if power == common.High {
		err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
		if err != nil {
			log.Logger.Error("Errjson:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		} else {
			if user.Profession == "学生" {
				user.Power = common.No
			} else {
				user.Power = common.Low
			}
			user.Password = "111111"
			_, err := models.UserService.CreateUser(&user)
			if err != nil {
				log.Logger.Error("ErrMysql", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			} else {

				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			}
		}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrPermission}
	}
	this.ServeJSON()
}

func (this *UserController) ModifyUserInfo() {
	var userInfo models.User

	power := this.GetSession(common.SessionPower)
	if power == common.High {
		err := json.Unmarshal(this.Ctx.Input.RequestBody, &userInfo)
		if err != nil {
			log.Logger.Error("Errjson:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		} else {
			_, err := models.UserService.CreateUser(&userInfo)
			if err != nil {
				log.Logger.Error("ErrMysql", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			} else {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			}
		}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrPermission}
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
