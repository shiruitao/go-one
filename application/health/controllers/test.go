package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/health/common"
	"github.com/shiruitao/go-one/application/health/log"
	"github.com/shiruitao/go-one/application/health/models"
)

type TestController struct {
	beego.Controller
}

func (this *TestController) Get() {
	var (
		id struct {
			//ID       uint32 `json:"id"`
			Answer1  int    `json:"answer1"`
			Answer2  int    `json:"answer2"`
			Answer3  int    `json:"answer3"`
			Answer4  int    `json:"answer4"`
			Answer5  int    `json:"answer5"`
			Answer6  int    `json:"answer6"`
			Answer7  int    `json:"answer7"`
			Answer8  int    `json:"answer8"`
			Answer9  int    `json:"answer9"`
			Answer10 int    `json:"answer10"`
		}
		score int
		res   string
		record models.Record
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &id)

	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		score = id.Answer1 + id.Answer2 + id.Answer3 + id.Answer4 + id.Answer5 + id.Answer6 + id.Answer7 + id.Answer8 + id.Answer9 + id.Answer10
		switch {
		case score >= 40 && score < 60:
			res = "健康状况已经达到令人担忧的状况了，最好不要承担繁重的工作，否则会使你的健康状况雪上加霜。平时应注意睡眠质量，多吃水果、蔬菜及蛋白质类食物，合理补充维生素及矿物质。"
		case score >= 60 && score < 80:
			res = "健康状况一般，容易疲倦，体力不支，不太适宜紧张而又强度大的工作，否则健康状况会下降。注意改善日常生活习惯，均衡饮食，适当补充维生素和矿物质。"
		case score >= 80 && score <= 100:
			res = "健康状况良好，精力充沛，充满活力，能够适应紧张而又强度大的工作。"
		default:
			res = "测试出错, 请答完全部题目!"

		}
		record.Name = this.GetSession(common.SessionUserName).(string)
		record.Result = res
		record.Score = score
		_, err := models.RecordService.Add(&record)
		log.Logger.Info("", err)
		this.Data["json"] = map[string]interface{}{"data": res, "score": score}
	}
	this.ServeJSON()
}
