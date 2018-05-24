package controllers

import (
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/community/common"
	"github.com/shiruitao/go-one/application/community/log"
	"github.com/shiruitao/go-one/application/community/models"
	"encoding/json"
)

type EconomicController struct {
	beego.Controller
}

func (this *EconomicController) GetEconomic() {
	var area struct{
		Area string `json:"area"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &area)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		economic, num, err := models.EconomicService.Get(area.Area)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "data": economic, "number": num}
		}
	}
	this.ServeJSON()
}
