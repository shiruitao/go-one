/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Inc.
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
 *     Initial: 2018/01/02        Shi Ruitao
 */
package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/shiruitao/go-one/application/cheng/log"
)

type MessageServiceProvider struct{}

var MessageService *MessageServiceProvider

type Message struct {
	Id      int       `orm:"column(id)"`
	Title   string    `orm:"column(title)"`
	Content string    `orm:"column(content)"`
	State   int       `orm:"column(state)"`
	Label1  string    `orm:"column(label1)"`
	Label2  string    `orm:"column(label2)"`
	Brief   string    `orm:"column(brief)"`
	Created time.Time `orm:"column(created)"`
}

// 注册模型
func init() {
	orm.RegisterModel(new(Message))
}

func (insert *MessageServiceProvider) Insert(content Message) error {
	content.State = 1
	o := orm.NewOrm()
	sql := "INSERT INTO blog.message(title, content, state, Label1, Label2, brief,created) VALUES (?, ?, ?, ?, ?, ?, ?)"
	values := []interface{}{content.Title, content.Content, content.State, content.Label1, content.Label2, content.Brief,time.Now().Format("2006-01-02 15:04:05")}
	_, err := o.Raw(sql, values).Exec()
	return err
}

type M struct {
	title   string
	content string
	label   string
}

func (readAll *MessageServiceProvider) ReadAll() ([]Message, int64, error) {
	var messages []Message
	o := orm.NewOrm()
	num, err := o.Raw("SELECT id, title, content, state,label1, label2, brief, created FROM message WHERE state = 1 ORDER BY id DESC").QueryRows(&messages)
	return messages, num, err
}

func (readLabel *MessageServiceProvider) ReadLabel(label1, label2 string) ([]Message, int64, error) {
	o := orm.NewOrm()
	var messages []Message
	num, err := o.Raw("SELECT id, title, content, label1, label2, brief, created FROM message WHERE label1 = ? OR label2 = ? AND state = 1 ORDER BY id DESC", label1, label2).QueryRows(&messages)
	return messages, num, err
}

// 大小写问题 ---------------------------------
//func regexp(s string) string {
//	var a string
//	capital := make([]string,len(s))
//	lower := make([]string,len(s))
//	for i,v:= range s {
//		lower[i]=strings.ToLower(string(v))
//		capital[i]=strings.ToUpper(string(v))
//	}
//
//	for i := 0; i < len(s); i++ {
//		b := fmt.Sprintf("[%s,%s]",lower[i],capital[i])
//		a = a+b
//	}
//
//	return a
//}
// ----------------------------------------------

func (readtitleContent *MessageServiceProvider) ReadTitleContent(title string) ([]Message, int64, error) {
	o := orm.NewOrm()
	var messages []Message

	//cond := orm.NewCondition()
	//cond1 := cond.Or("title__icontains", title).Or("content__icontains", title).And("state__exact", 1)
	//qs := o.QueryTable("message")
	//num, err := qs.SetCond(cond1).All(&messages)

	//qs := o.QueryTable("message")
	//num, err := qs.Filter("title__iexact", title).All(&messages)
	title = "%" + title + "%"
	num, err := o.Raw("SELECT id, title, content, label1, label2, brief, created FROM message WHERE title LIKE ? OR content LIKE ? AND state = 1 ORDER BY id DESC", title, title).QueryRows(&messages)
	return messages, num, err
}

func (readTime *MessageServiceProvider) ReadTime(date string) ([]Message, int64, error) {
	o := orm.NewOrm()
	var messages []Message
	fmt.Println(date)
	date = date + "%"
	fmt.Println(date)
	num, err := o.Raw("SELECT id, title, content, label1, label2, brief, created FROM message WHERE created LIKE ? AND state = 1 ORDER BY id DESC", date).QueryRows(&messages)
	return messages, num, err
}

func (del *MessageServiceProvider) Delete(id int) error {
	o := orm.NewOrm()
	_, err := o.Raw("UPDATE message SET state = 0 WHERE id = ?", id).Exec()
	return err
}

func (delTest *MessageServiceProvider) DeleteTest(id int) (int64, error) {
	o := orm.NewOrm()
	res, err := o.Raw("DELETE FROM message WHERE id = ?", id).Exec()
	num, _ := res.RowsAffected()
	if err != nil {
		log.Logger.Error("Delete:", err)
	}
	return num, err
}
