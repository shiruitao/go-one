package controllers

import (
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/hjjc/common"
	"github.com/shiruitao/go-one/application/hjjc/log"
	"github.com/shiruitao/go-one/application/hjjc/models"
)

type AnalysesController struct {
	beego.Controller
}

func (this *AnalysesController) Average() {
	info, err := models.HistoryService.Get()
	if err != nil {
		log.Logger.Error("ErrMysql", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{"data": info}
	}
	this.ServeJSON()

}
