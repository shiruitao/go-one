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

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/shiruitao/go-one/application/bbs/initorm"
)

type ArticleServiceProvider struct{}

var (
	ArticleService *ArticleServiceProvider
)

type (
	// Article represents the article information.
	Article struct {
		Id       bson.ObjectId `bson:"_id,omitempty"`
		Title    string        `bson:"title"`
		Content  string        `bson:"content"`
		AuthorID uint32        `bson:"authorID"`
		Created  string        `bson:"created"`
		Image1   string        `bson:"image1"`
		Image2   string        `bson:"image2"`
		Image3   string        `bson:"image3"`
		Video    string        `bson:"video"`
		Active   bool          `bson:"active"`
	}
	Connection struct {
		session    *mgo.Session
		collection *mgo.Collection
		Database   string
		Name       string
	}
)

func CollectionArticle() initorm.Mongodb {
	m := initorm.ConnectMongo("article")
	m.C.EnsureIndex(mgo.Index{
		Key:        []string{"Title"},
		Unique:     false,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	})
	return m
}

func (this *ArticleServiceProvider) New(a *Article) (string, error) {
	m := CollectionArticle()
	defer m.S.Close()

	// 生成 ObjectId
	a.Id = bson.NewObjectId()

	a.Created = time.Now().Format("2006-01-02 15:04:05")
	a.Active = true
	err := m.C.Insert(a)
	if err != nil {
		return "", err
	}

	return a.Id.Hex(), nil
}
