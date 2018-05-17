package models

import (
	"time"
	
	"github.com/astaxie/beego/orm"
)

type UserServiceProvider struct{}

var UserService *UserServiceProvider

type User struct {
	ID      uint32    `orm:"column(id);pk;auto"`
	Name    string    `orm:"column(name);null;utf8_bin" json:"name"`
	Avatar  string    `orm:"column(avatar)" json:"avatar"`
	IsAdmin bool      `orm:"column(isadmin)" json:"is_admin"`
	Created time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (this *UserServiceProvider) CreateUser(u *User) (uint32, bool, error) {
	var (
		user User
		id   int64
		err  error
	)
	user.Name = u.Name
	user.IsAdmin = false
	o := orm.NewOrm()
	if _, id, err = o.ReadOrCreate(&user, "unionid"); err != nil {
		return uint32(id), false, err
	}
	user = User{ID: uint32(id)}
	err = o.Read(&user)
	return uint32(id), user.IsAdmin, err
}
