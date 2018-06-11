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

func (this *CommentController) Add()  {
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

func (this *CommentController) Get() {

}

