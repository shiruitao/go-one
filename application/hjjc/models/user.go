package models

import (

	"github.com/astaxie/beego/orm"
)

type UserServiceProvider struct{}

var UserService *UserServiceProvider

type User struct {
	ID       uint32    `orm:"column(id);pk;auto"`
	Name     string    `orm:"column(name)" json:"name"`
	RealName string    `orm:"column(realname)" json:"real_name"`
	Password string    `orm:"column(password)" json:"password"`
	UserRole string    `orm:"column(userrole)" json:"user_role"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (this *UserServiceProvider) CreateUser(userInfo *User) (int64, error) {
	var (
		user User
		id   int64
		err  error
	)
	user.Name = userInfo.Name
	user.RealName = userInfo.RealName
	user.Password = userInfo.Password
	user.UserRole = userInfo.UserRole

	o := orm.NewOrm()

	id, err = o.Insert(&user)
	if err != nil {
		return id, err
	}

	return id, nil
}
