package models

import (
	"github.com/astaxie/beego/orm"
)

type UserServiceProvider struct{}

var UserService *UserServiceProvider

type (
	User struct {
		ID      uint32    `orm:"column(id);pk;auto"`
		Name    string    `orm:"column(name)" json:"name"`
		Sex     string    `orm:"column(sex)" json:"sex"`
		Weight  int       `orm:"column(weight)" json:"weight"`
		High    int       `orm:"column(high)" json:"high"`
		Age     int       `orm:"column(age)" json:"age"`
	}
)

func init() {
	orm.RegisterModel(new(User))
}

func (this *UserServiceProvider) CreateUser(userInfo *User) (int64, error) {
	var (
		user User
	)
	user.Name = userInfo.Name
	user.High = userInfo.High
	user.Weight = userInfo.Weight
	user.Sex = userInfo.Sex
	user.Age = userInfo.Age

	o := orm.NewOrm()

	return o.Insert(&user)
}

func (this *UserServiceProvider) Login(name string) (uint32, error) {
	var (
		user User
	)

	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("name", name).One(&user)

	return user.ID, err
}
