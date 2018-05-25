package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/community/log"
	"github.com/shiruitao/go-one/application/community/common"
	"github.com/shiruitao/go-one/application/community/models"
)

type NoticeController struct {
	beego.Controller
}

func (this *NoticeController) Create() {
	var (
		h models.Notice
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &h)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		_, err := models.NoticeService.Create(&h)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *NoticeController) GetAll() {
	notice, err := models.HouseService.GetAll()
	if err != nil {
		log.Logger.Error("ErrMysql", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": notice}
	}
	this.ServeJSON()
}
