package models

import (
	"time"
	"errors"

	"github.com/astaxie/beego/orm"

	"github.com/shiruitao/go-one/application/bbs/utility"
	"github.com/shiruitao/go-one/application/bbs/common"
)

type UserServiceProvider struct{}

var UserService *UserServiceProvider

type User struct {
	ID        uint32    `orm:"column(id);pk;auto"`
	Name      string    `orm:"column(name);not null;utf8_bin" json:"name"`
	Nickname  string    `orm:"column(nickname);not null;utf8_bin" json:"nickname"`
	Password  string    `orm:"column(password)";type:varchar(128)" json:"password"`
	Avatar    string    `orm:"column(avatar)" json:"avatar"`
	IsAdmin   bool      `orm:"column(isadmin)" json:"is_admin"`
	LastLogin time.Time `orm:"column(lastlogin);type(datetime)"`
	Created   time.Time `orm:"column(created);auto_now_add;type(datetime)"`
	Isactive  bool      `orm:"column(isactive)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (this *UserServiceProvider) CreateUser(name, nickname, password, avatar, confirmPassword string) (int64, error) {
	var (
		user User
		id   int64
		err  error
	)
	user.Isactive = true
	user.Name = name
	user.Nickname = nickname
	user.Avatar = avatar
	user.IsAdmin = false
	user.LastLogin = time.Now()

	if len(password) < 6 || len(password) > 30 || password != confirmPassword {
		err = errors.New(common.ErrWrongPass)
		return id, err
	}

	hash, err := utility.GenerateHash(password)
	user.Password = string(hash)
	if err != nil {
		return id, err
	}
	o := orm.NewOrm()

	id, err = o.Insert(&user)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (this *UserServiceProvider) Login(name, password string) (bool, bool, error) {
	var (
		user User
	)

	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("name", name).One(&user)
	if err != nil {
		return false, false, err
	} else if !utility.CompareHash([]byte(user.Password), password) {
		return false, false, errors.New(common.ErrWrongPass)
	}

	user.LastLogin = time.Now()
	_, err = o.Update(&user, "lastlogin")
	return user.IsAdmin, true, nil
}

func (this *UserServiceProvider) ChangePassword(name, newPassword string) error {
	var (
		user User
	)

	o := orm.NewOrm()
	hash, err := utility.GenerateHash(newPassword)
	if err != nil {
		return err
	}
	newPassword = string(hash)

	user = User{ID: 1}
	user.Password = newPassword
	_, err = o.Update(&user, "password")
	return err
}

//func (this *UserServiceProvider) ChangeInfo() error {
//
//	return
//}
