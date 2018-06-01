/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/05/19        Shi Ruitao
 */

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type ArticleServiceProvider struct{}

var (
	ArticleService *ArticleServiceProvider
)

type (
	Article struct {
		ID       uint32    `orm:"column(id);pk;auto"`
		Title    string    `orm:"column(title)" json:"title"`
		Content  string    `orm:"column(content)" json:"content"`
		AuthorID uint32    `orm:"column(Authorid)"`
		Image1   string    `orm:"column(image1)" json:"image1"`
		Image2   string    `orm:"column(image2)" json:"image2"`
		Image3   string    `orm:"column(image3)" json:"image3"`
		Video    string    `orm:"column(video)" json:"video"`
		IsActive bool      `orm:"column(isactive)"`
		Created  time.Time `orm:"column(created)"`
	}
)

func init() {
	orm.RegisterModel(new(Article))
}

func (*ArticleServiceProvider) New(info *Article) error {
	var (
		article Article
	)
	article.Title = info.Title
	article.Content = info.Content
	article.AuthorID = info.AuthorID
	article.Image1 = info.Image1
	article.Image2 = info.Image2
	article.Image3 = info.Image3
	article.Video = info.Video
	article.IsActive = true

	o := orm.NewOrm()
	_, err := o.Insert(&article)

	return err
}

func (*ArticleServiceProvider) Get() (*[]Article, error) {
	var (
		article []Article
	)

	o := orm.NewOrm()
	_, err := o.Raw("SELECT * FROM article").QueryRows(&article)

	return &article, err
}
