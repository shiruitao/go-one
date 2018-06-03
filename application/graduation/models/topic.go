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

	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/shiruitao/go-one/application/graduation/common"
)

type TopicServiceProvider struct{}

var TopicService *TopicServiceProvider

type Topic struct {
	ID          uint32    `orm:"column(id);pk;auto"`
	Name        string    `orm:"column(name)" json:"title"`
	TeacherID   uint32    `orm:"column(teacherid)" json:"teacher_id"`
	TeacherName string    `orm:"column(teachername)" json:"teacher_name"`
	Type        int8      `orm:"column(type)" json:"type"` // 1 -> 教室发布,管理员未审核,学生无权浏览; 2 -> 管理审核通过,学生可选; 3 -> 学生选定,等待教室确认; 4 -> 教室确认,最终状态
	StudentID   uint32    `orm:"column(studentid)"`
	StuName     string    `orm:"column(studentname)"`
	StuNum      string    `orm:"column(studentnum)"`
	Created     time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Topic))
}

func (*TopicServiceProvider) Create(info *Topic) (int64, error) {
	var (
		topic Topic
	)

	topic.Name = info.Name
	topic.TeacherID = info.TeacherID
	topic.Type = common.CanSelect
	topic.TeacherName = info.TeacherName

	o := orm.NewOrm()

	return o.Insert(&topic)
}

func (*TopicServiceProvider) Select(id, userID uint32, stuName string, stuNum string) (int64, error) {
	var (
		topic Topic
	)
	o := orm.NewOrm()

	t := get(id, o)
	if t.Type != common.CanSelect {
		return 0, errors.New("No choice allowed!")
	}

	topic.ID = id
	topic.Type = common.Selected
	topic.StudentID = userID
	topic.StuName = stuName
	topic.StuNum = stuNum

	return o.Update(&topic, "studentid","type", "stuname", "stunum")
}

func (*TopicServiceProvider) Duplicate(userID uint32) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	cnt, err := qs.Filter("studentid", userID).Count()
	return cnt, err
}

func (*TopicServiceProvider) Bake(id uint32) (int64, error) {
	var (
		topic Topic
	)
	o := orm.NewOrm()

	t := get(id, o)
	if t.Type == common.Finish {
		return 0, errors.New("No choice allowed!")
	}
	topic.ID = id
	topic.StuNum = ""
	topic.StuName = ""
	topic.StudentID = 0
	topic.Type = common.CanSelect
	return o.Update(&topic, "type", "studentid", "studentname", "studentnum")
}

func (*TopicServiceProvider) StudentGetTopic(user string) (*[]Topic, Topic, error) {
	var (
		topic       []Topic
		topicSelect Topic
	)

	o := orm.NewOrm()

	qs := o.QueryTable("topic")
	qs.Filter("studentnum", user).One(&topicSelect)
	_, err := qs.Exclude("type", common.Affirm).All(&topic)

	return &topic, topicSelect, err
}

func (*TopicServiceProvider) TeacherGetTopic(id uint32) (*[]Topic, error) {
	var (
		topic []Topic
	)

	o := orm.NewOrm()
	qs := o.QueryTable("topic")
	_, err := qs.Filter("teacherid", id).All(&topic)
	return &topic, err
}

func (*TopicServiceProvider) AdminGetTopic() (*[]Topic, error) {
	var (
		topic []Topic
	)
	o := orm.NewOrm()
	err := o.Read(&topic)
	return &topic, err
}

func (*TopicServiceProvider) AdminCheck(id uint32) error {
	o := orm.NewOrm()
	top := get(id, o)
	if top.Type != 1 {
		return errors.New("No choice allowed!")
	}
	topic := Topic{
		ID:   id,
		Type: common.CanSelect,
	}

	_, err := o.Update(&topic, "type")

	return err
}

func (*TopicServiceProvider) TeacherModifyTopic(t *Topic) error {
	var (
		topic Topic
	)
	o := orm.NewOrm()
	top := get(t.ID, o)
	if top.Type == common.Selected || top.Type == common.Finish {
		return errors.New("No choice allowed!")
	}
	topic.ID = t.ID
	topic.Name = t.Name
	topic.Type = common.Affirm
	_, err := o.Update(&topic, "name", "type")
	return err
}

func (*TopicServiceProvider) TeacherVerify(id uint32) error {
	o := orm.NewOrm()
	t := get(id, o)
	if t.Type != common.Selected {
		return errors.New("No choice allowed!")
	}

	topic := Topic{
		ID:   id,
		Type: common.Finish,
	}

	_, err := o.Update(&topic, "type")
	return err
}

func get(id uint32, o orm.Ormer) Topic {
	topic := Topic{ID: id}
	o.Read(&topic)
	return topic
}
