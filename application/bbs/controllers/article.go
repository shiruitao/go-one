package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"

	"github.com/shiruitao/go-one/application/bbs/common"
	"github.com/shiruitao/go-one/application/bbs/log"
	"github.com/shiruitao/go-one/application/bbs/models"
)

type ArticleController struct {
	beego.Controller
}

type (
	createArticle struct {
		Title    string `json:"title"`
		Content  string `json:"content"`
		AuthorID uint32 `json:"authorid"`
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
		article.AuthorID = this.GetSession(common.SessionUserID).(uint32)
		article.Image1 = addArticle.Image1
		article.Image2 = addArticle.Image2
		article.Image3 = addArticle.Image3
		article.Video = addArticle.Video

		_, err = models.ArticleService.New(&article)
	}
}
