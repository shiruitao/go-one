package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"github.com/shiruitao/go-one/application/cheng/log"
)

type MessageServiceProvider struct{}

var MessageService *MessageServiceProvider

type Message struct {
	Id      int    `orm:"column(id)"`
	Title   string `orm:"column(title)"`
	Content string `orm:"column(content)"`
	State   int    `orm:"column(state)"`
	Label   string `orm:"column(label)"`
}

// 注册模型
func init() {
	orm.RegisterModel(new(Message))
}

func (insert *MessageServiceProvider) Insert(content Message) (int64, error) {
	fmt.Println("model中", content)
	content.State = 1
	o := orm.NewOrm() // 创建一个 Ormer
	// NewOrm 的同时会执行 orm.BootStrap (整个 app 只执行一次)，用以验证模型之间的定义并缓存。
	// o.Using("blog") // 默认使用 default，你可以指定为其他数据库
	id, err := o.Insert(&content)
	if err != nil {
		return 0, err
	}
	return id, err
}


func (readLabel *MessageServiceProvider) ReadLabel(label string) []Message {
	o := orm.NewOrm()
	var messages []Message
	num, err := o.Raw("SELECT * FROM message where label = ? and state = 1", label).QueryRows(&messages)
	if err == nil {
		fmt.Println("message content: ", messages, num)
	}
	return messages
}

func (readtitleContent *MessageServiceProvider) ReadTitleContent(title string) ([]*Message, int64) {
	var message []*Message
	o := orm.NewOrm()
	qs := o.QueryTable("message")
	num, err := qs.Filter("title__icontains", title).All(&message)
	if err != nil {
		log.Logger.Error("qs.Filter:", err)
	}
	return message, num
}

func (del *MessageServiceProvider) Delete(id int) (int64, error) {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE message SET state = 0 WHERE id = ?", id).Exec()
	num, _ := res.RowsAffected()
	if err != nil {
		log.Logger.Error("Delete:", err)
	}
	return num, err

}

func (delTest *MessageServiceProvider) DeleteTest(id int) (int64, error) {
	o := orm.NewOrm()
	res, err := o.Raw("DELETE from message WHERE id = ?", id).Exec()
	num, _ := res.RowsAffected()
	if err != nil {
		log.Logger.Error("Delete:", err)
	}
	return num, err
}

func (up *MessageServiceProvider) Update(id int) {
	o := orm.NewOrm()
	message := Message{Id: id}
	if o.Read(&message) == nil {
		message.Title = "MyName"
		if num, err := o.Update(&message); err == nil {
			fmt.Println(num)
		}
	} else {
		fmt.Println("Id:", id, "不存在")
	}
}
