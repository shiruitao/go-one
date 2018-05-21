package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/bbs/common"
	"github.com/shiruitao/go-one/application/bbs/log"
	"github.com/shiruitao/go-one/application/bbs/models"
	"fmt"
)

type ArticleController struct {
	beego.Controller
}

type (
	createArticle struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		//AuthorID uint32 `json:"author_id"`
		Image1   string `json:"image1"`
		Image2   string `json:"image2"`
		Image3   string `json:"image3"`
		Video    string `json:"video"`
	}
)

func (this *ArticleController) ArticleCreate() {
	var (
		addArticle createArticle
		article    models.Article
	)

	err := json.Unmarshal(this.Ctx.Input.RequestBody, &addArticle)
	if err != nil {
		log.Logger.Error("Errjson:", err)
		this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrInvalidParam}
	} else {
		article.Title = addArticle.Title
		article.Content = addArticle.Content
		fmt.Println("uid: ", this.GetSession(common.SessionUserID), this.GetSession(common.SessionUserID).(uint32))
		article.AuthorID = this.GetSession(common.SessionUserID).(uint32)
		article.Image1 = addArticle.Image1
		article.Image2 = addArticle.Image2
		article.Image3 = addArticle.Image3
		article.Video = addArticle.Video

		_, err = models.ArticleService.New(&article)
		if err != nil {
			log.Logger.Error("ErrDB:", err)
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrMysqlQuery}
		} else {
			this.Data["json"] = map[string]interface{}{common.RespKeyStatus: common.ErrSucceed}
		}

	}
	this.ServeJSON()
}
