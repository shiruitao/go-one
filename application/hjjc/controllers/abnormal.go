package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/hjjc/models"
)

type AbnormalController struct {
	beego.Controller
}

func (this *AbnormalController) Get() {
	info, _ := models.DevicesService.Get()
	historyData, _ := models.HistoryService.Get()

	this.Data["json"] = map[string]interface{}{"devices": info, "historyData": historyData}
	this.ServeJSON()
}
