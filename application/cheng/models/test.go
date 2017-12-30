package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
)

func Helloworld(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
func Add(a, b int) int {
	c := a + b

	return c
}

type MessageServiceProvider struct{}

var MessageService *MessageServiceProvider

type Message struct {
	Id      int
	Title   string
	Content string
}
// 注册模型
func init() {
	orm.RegisterModel(new(Message))
}

func (insert *MessageServiceProvider) Insert() (int64, error) {
	o := orm.NewOrm() // 创建一个 Ormer
	// NewOrm 的同时会执行 orm.BootStrap (整个 app 只执行一次)，用以验证模型之间的定义并缓存。
	o.Using("blog") // 默认使用 default，你可以指定为其他数据库
	message := new(Message)
	message.Title = "my first bolg"
	message.Content = "默认使用 default，你可以指定为其他数据库"
	id, err := o.Insert(message)
	if err != nil {
		return 0, err
	}

	return id, err
}
func (read *MessageServiceProvider) Read() error {
	o := orm.NewOrm()
	message := Message{Id: 1000}
	
	err := o.Read(&message)
	
	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	} else {
		fmt.Println(message.Id, message.Title)
	}
	return err
}
