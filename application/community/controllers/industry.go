package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/community/log"
	"github.com/shiruitao/go-one/application/community/common"
	"github.com/shiruitao/go-one/application/community/models"
)

type IndustryController struct {
	beego.Controller
}

func (this *IndustryController) GetIndustry() {
	var area struct{
		Area string `json:"area"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &area)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		industry, num, err := models.IndustryService.Get(area.Area)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": industry, "number": num}
		}
	}
	this.ServeJSON()
}
