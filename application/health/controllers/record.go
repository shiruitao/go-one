package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/health/models"
	"github.com/shiruitao/go-one/application/health/common"
)

type RecordController struct {
	beego.Controller
}

func (this *RecordController) Get() {
	name := this.GetSession(common.SessionUserName).(string)
	record := models.RecordService.Get(name)
	this.Data["json"] = record
	this.ServeJSON()
}
