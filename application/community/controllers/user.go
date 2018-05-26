package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/community/common"
	"github.com/shiruitao/go-one/application/community/log"
	"github.com/shiruitao/go-one/application/community/models"
)

type UserController struct {
	beego.Controller
}

func (this *UserController) CreateUser() {
	var user models.User

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		_, err := models.UserService.CreateUser(&user)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *UserController) UserUpdate() {
	var user models.User

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &user)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		err := models.UserService.ChangeInfo(&user)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}
	}
	this.ServeJSON()
}

func (this *UserController) GetUser() {
	user, num, err := models.UserService.GetUser()
	if err != nil {
		log.Logger.Error("ErrMysql", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "number": num, "info": user}
	}
	this.ServeJSON()
}

func (this *UserController) GetUserByName() {
	var name struct {
		Name string `json:"name"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &name)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		user, num, err := models.UserService.GetUserByName(name.Name)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "number": num, "info": user}
		}
		this.ServeJSON()
	}
}

func (this *UserController) GetUserByArea() {
	var area struct {
		Area string `json:"area"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &area)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		user, num, err := models.UserService.GetUserByArea(area.Area)
		if err != nil {
			log.Logger.Error("ErrMysql", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "number": num, "info": user}
		}
		this.ServeJSON()
	}
}

func (this *UserController) GetUserByAge() {
	var area struct {
		Area string `json:"area"`
	}

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &area)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	}
	num1, num2, num3, num4, num5, err := models.UserService.GetUserByAge(area.Area)
	if err != nil {
		log.Logger.Error("ErrMysql", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
	} else {
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed, "number1": num1, "number2": num2, "number3": num3, "number4": num4, "number5": num5}
	}
	this.ServeJSON()
}

func (this *UserController) Weather() {
	type (
		b struct {
			Date string `json:"date"`
			High string `json:"high"`
			Low  string `json:"low"`
			Day string `json:"day"`
			Text string `json:"text"`
			Wind string `json:"wind"`
		}
		a struct {
			Future []b `json:"future"`
		}

	)
	var weather struct {
		Weather []a    `json:"weather"`
	}
	resp, _ := http.Get("http://tj.nineton.cn/Heart/index/all?city=CHHE010000&language=zh-chs&unit=c&aqi=city&alarm=1&key=78928e706123c1a8f1766f062bc8676b")
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(body, &weather)
	this.Data["json"] =  weather
	this.ServeJSON()
}
