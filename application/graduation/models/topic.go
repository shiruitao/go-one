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
 *     Initial: 2018/05/25        Shi Ruitao
 */

package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/shiruitao/go-one/application/graduation/common"
)

type TopicServiceProvider struct{}

var TopicService *TopicServiceProvider

type Topic struct {
	ID        uint32    `orm:"column(id);pk;auto"`
	Name      string    `orm:"column(name)" json:"name"`
	TeacherID uint32    `orm:"column(teacherid)" json:"teacher_id"`
	Type      int8      `orm:"column(type)"`
	StuName   string    `orm:"column(stuname)" json:"stu_name"`
	StuNum    string    `orm:"column(stunum)" json:"stu_num"`
	Created   time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (*TopicServiceProvider) Create(info *Topic) (int64, error) {
	var (
		topic Topic
	)

	topic.Name = info.Name
	topic.TeacherID = info.TeacherID
	topic.Type = common.CanSelect

	o := orm.NewOrm()

	return o.Insert(&topic)
}

func (*TopicServiceProvider) Select(id uint32, stuName string, stuNum string) (int64, error) {
	var (
		topic Topic
	)
	topic = Topic{ID: id}
	topic.Type = common.Selected
	topic.StuName = stuName
	topic.StuNum = stuNum

	o := orm.NewOrm()

	return o.Update(&topic, "type", "stuid", "stuname", "stunum")
}

func (*TopicServiceProvider) Bake(id uint32) (int64, error) {
	var (
		topic Topic
	)
	topic = Topic{ID: id}
	topic.Type = common.CanSelect

	o := orm.NewOrm()
	return o.Update(&topic, "type")
}

func (*TopicServiceProvider) GetType(id uint32) int8 {
	var (
		topic Topic
	)
	topic = Topic{ID: id}
	o := orm.NewOrm()
	o.Read(&topic)
	return topic.Type
}