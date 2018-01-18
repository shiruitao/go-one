package models

import (
	"github.com/astaxie/beego/orm"
)

type UserServerProvider struct {
}

var UserServer *UserServerProvider = &UserServerProvider{}

type User struct {
	Id int
	Name  string `form:"name"`
	Age   int    `form:"age"`
	Email string
	Sex   string
}

func init() {
	orm.RegisterModel(new(User))
}

func (create *UserServerProvider) Create(info User) error {
	o := orm.NewOrm()
	sql := "INSERT INTO luoo.user(name, age, email, sex) VALUES (?, ?, ?)"
	values := []interface{}{info.Name, info.Age, info.Email, info.Sex}
	raw := o.Raw(sql, values)
	_, err := raw.Exec()
	return err
}
