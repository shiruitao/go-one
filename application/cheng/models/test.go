package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

type MessageServiceProvider struct{}

var MessageService *MessageServiceProvider

type Message struct {
	Id      int
	Title   string
	Content string
	State   int
	Label   string
}
// 注册模型
func init() {
	orm.RegisterModel(new(Message))
}

func (insert *MessageServiceProvider) Insert(content Message) (int64, error) {
	fmt.Println("model中", content)
	o := orm.NewOrm() // 创建一个 Ormer
	// NewOrm 的同时会执行 orm.BootStrap (整个 app 只执行一次)，用以验证模型之间的定义并缓存。
	// o.Using("blog") // 默认使用 default，你可以指定为其他数据库
	id, err := o.Insert(&content)
	if err != nil {
		return 0, err
	}
	return id, err
}

func (read *MessageServiceProvider) Read(id int) error {
	o := orm.NewOrm()
	message := Message{Id: id}
	fmt.Println("输出结果:", message)
	err := o.Read(&message)
	fmt.Println("err", err)
	fmt.Println("err:", err)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(message)
	}
	return err
}

func (read *MessageServiceProvider) ReadLabel(label string) Message{
	o := orm.NewOrm()
	var messages Message
	num, err := o.Raw("SELECT * FROM message where label = ?", label).QueryRows(&messages)
	if err == nil {
		fmt.Println("user nums: ", num)
		fmt.Println("user nums: ", messages)
	}
	return messages
}


func (read *MessageServiceProvider) Read_title_content(label string) error {
	o := orm.NewOrm()
	message := Message{Label: label}
	fmt.Println("输出结果:", message)
	err := o.Read(&message)
	fmt.Println("err", err)
	fmt.Println("err:", err)
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(message.Id, message.Title)
	}
	return err
}

func (del *MessageServiceProvider) Delete(id int) (int64, error) {
	fmt.Println("model中:", id)
	o := orm.NewOrm()
	num, err := o.Delete(&Message{Id: id})
	if err == nil {
		fmt.Println(num)
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
