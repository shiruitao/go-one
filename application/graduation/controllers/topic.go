package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/graduation/common"
	"github.com/shiruitao/go-one/application/graduation/log"
	"github.com/shiruitao/go-one/application/graduation/models"
)

type TopicController struct {
	beego.Controller
}

func (this *TopicController) CreateTopic() {
	var topic models.Topic

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &topic)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		_, err := models.TopicService.Create(&topic)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *TopicController) Select() {
	var id struct {
		ID uint32 `json:"id"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		userID := this.GetSession(common.SessionUserID).(uint32)
		user := models.UserService.GetUserByID(userID)
		_, err := models.TopicService.Select(id.ID, user.Name, user.Number)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}
