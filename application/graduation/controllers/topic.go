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
	var (
		topic models.Topic
		title struct {
			Name string `json:"name"`
		}
	)

	power := this.GetSession(common.SessionPower).(int8)
	if power != 2 {
		log.Logger.Error("Have no legal power")
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrPermission}
	} else {
		err := json.Unmarshal(this.Ctx.Input.RequestBody, &title)
		if err != nil {
			log.Logger.Error("Errjson:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		} else {
			userID := this.GetSession(common.SessionUserID).(uint32)
			userName := this.GetSession(common.SessionUserName).(string)
			topic.Name = title.Name
			topic.TeacherName = userName
			topic.TeacherID = userID
			topic.StudentID = 0
			_, err := models.TopicService.Create(&topic)
			if err != nil {
				log.Logger.Error("ErrMysql", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			} else {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			}
		}
	}
	this.ServeJSON()
}

func (this *TopicController) Select() {
	var id struct {
		ID uint32 `json:"id"`
	}

	userID := this.GetSession(common.SessionUserID).(uint32)

	cnt, _ := models.TopicService.Duplicate(userID)
	if cnt == 1 {
		log.Logger.Error("Duplicate select")
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: "Duplicate select"}
	} else {
		err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
		if err != nil {
			log.Logger.Error("Errjson:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
		} else {
			userName := this.GetSession(common.SessionUserName).(string)
			userNumber := this.GetSession(common.SessionUserNum).(string)

			_, err := models.TopicService.Select(id.ID, userID, userName, userNumber)
			if err != nil {
				log.Logger.Error("ErrMysql", err)
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
			} else {
				this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
			}
		}

	}

	this.ServeJSON()
}

func (this *TopicController) Back() {
	var id struct {
		ID uint32 `json:"id"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		_, err = models.TopicService.Bake(id.ID)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *TopicController) StudentGetTopic() {
	userNumber := this.GetSession(common.SessionUserNum).(string)
	topic, topicSelect, err := models.TopicService.StudentGetTopic(userNumber)
	if err != nil {
		log.Logger.Error("ErrMysql", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": topic, "select": topicSelect}
	}
	this.ServeJSON()
}

func (this *TopicController) TeacherGetTopic() {
	id := this.GetSession(common.SessionUserID).(uint32)

	topic, err := models.TopicService.TeacherGetTopic(id)
	if err != nil {
		log.Logger.Error("ErrMysql", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": topic}
	}
	this.ServeJSON()
}

func (this *TopicController) AdminGetTopic() {
	topic, err := models.TopicService.AdminGetTopic()
	if err != nil {
		log.Logger.Error("ErrMysql", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": topic}
	}
	this.ServeJSON()
}

func (this *TopicController) AdminCheck() {
	var id struct {
		ID uint32 `json:"id"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.TopicService.AdminCheck(id.ID)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *TopicController) TeacherVerify() {
	var id struct {
		ID   uint32 `json:"id"`
		Type int8   `json:"tpye"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.TopicService.TeacherVerify(id.ID)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *TopicController) TeacherModifyTopic() {
	var topic models.Topic
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &topic)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.TopicService.TeacherModifyTopic(&topic)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}
