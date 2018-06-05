package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/shiruitao/go-one/application/hjjc/common"
	"github.com/shiruitao/go-one/application/hjjc/log"
	"github.com/shiruitao/go-one/application/hjjc/models"
	"time"
)

type AnalysesController struct {
	beego.Controller
}

func (this *AnalysesController) Average() {
	var (
		average []float32
		t []time.Time
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &t)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		info, err := models.HistoryService.Get()
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			for i, value := range info {
				average[i] = (value.CurrentA + value.CurrentB + value.CurrentC)/3
				t[i] = value.DataTimeHour
			}
			this.Data["json"] = map[string]interface{}{"hour": t, "data": average}
		}
	}
	this.ServeJSON()

}
