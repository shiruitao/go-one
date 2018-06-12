package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/bbs/models"
	"encoding/json"
	"github.com/shiruitao/go-one/application/bbs/log"
	"github.com/shiruitao/go-one/application/bbs/common"
)

type CommentController struct {
	beego.Controller
}

func (this *CommentController) AddRep()  {
	var comment models.Comment

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &comment)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		userID := this.GetSession(common.SessionUserID).(uint32)
		userName := this.GetSession(comment.Replied).(string)
		comment.RepliedID = userID
		comment.Replied = userName
		err := models.CommentService.Add(&comment)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
}

func (this *CommentController) AddCreator()  {
	var comment models.Comment

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &comment)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		userID := this.GetSession(common.SessionUserID).(uint32)
		userName := this.GetSession(comment.Replied).(string)
		comment.CreatorID = userID
		comment.Creator = userName
		err := models.CommentService.Add(&comment)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
}

func (this *CommentController) Get() {
	var id struct {
		ID uint32 `json:"id"`
	}
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		c, err := models.CommentService.Get(id.ID)
		if err != nil {
			log.Logger.Error("ErrDB:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": c}
		}
	}
	this.ServeJSON()
}

