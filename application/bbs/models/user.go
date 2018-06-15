package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/shiruitao/go-one/application/bbs/common"
	"github.com/shiruitao/go-one/application/bbs/utility"
)

type UserServiceProvider struct{}

var UserService *UserServiceProvider

type (
	User struct {
		ID        uint32    `orm:"column(id);pk;auto"`
		Name      string    `orm:"column(name)" json:"name"`
		Nickname  string    `orm:"column(nickname)" json:"nickname"`
		Sex       string    `orm:"column(sex)" json:"sex"`
		RealName  string    `orm:"column(realname)" json:"real_name"`
		School    string    `orm:"column(school)" json:"school"`
		Password  string    `orm:"column(password)" json:"password"`
		Avatar    string    `orm:"column(avatar)" json:"avatar"`
		IsAdmin   bool      `orm:"column(isadmin)" json:"is_admin"`
		LastLogin time.Time `orm:"column(lastlogin);"`
		Created   time.Time `orm:"column(created);auto_now_add;type(datetime)"`
		Isactive  bool      `orm:"column(isactive)"`
	}
	UserInfo struct {
		Name            string `json:"name"`
		Nickname        string `json:"nickname"`
		Password        string `json:"password"`
		Sex             string `json:"sex"`
		RealName        string `json:"real_name"`
		School          string `json:"school"`
		Avatar          string `json:"avatar"`
		ConfirmPassword string `json:"confirm_password"`
	}
)

func init() {
	orm.RegisterModel(new(User))
}

func (this *UserServiceProvider) CreateUser(userInfo *UserInfo) (int64, error) {
	var (
		user User
		id   int64
		err  error
	)
	user.Isactive = true
	user.Name = userInfo.Name
	user.Nickname = userInfo.Nickname
	user.Avatar = userInfo.Avatar
	user.Sex = userInfo.Sex
	user.School = userInfo.School
	user.RealName = userInfo.RealName
	user.IsAdmin = false
	user.LastLogin = time.Now()

	if len(userInfo.Password) < 6 || len(userInfo.Password) > 30 || userInfo.Password != userInfo.ConfirmPassword {
		err = errors.New(common.ErrWrongPass)
		return id, err
	}

	hash, err := utility.GenerateHash(userInfo.Password)
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

func (this *UserServiceProvider) Login(name, password string) (User, bool, error) {
	var (
		user User
	)

	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("name", name).One(&user)
	if err != nil {
		return user, false, err
	} else if !utility.CompareHash([]byte(user.Password), password) {
		return user, false, errors.New(common.ErrWrongPass)
	}

	user.LastLogin = time.Now()
	_, err = o.Update(&user, "lastlogin")
	return user, true, nil
}

func (this *UserServiceProvider) ChangePassword(id uint32, newPassword string) error {
	var (
		user User
	)

	o := orm.NewOrm()
	hash, err := utility.GenerateHash(newPassword)
	if err != nil {
		return err
	}
	newPassword = string(hash)

	user = User{ID: id}
	user.Password = newPassword
	_, err = o.Update(&user, "password")
	return err
}

func (this *UserServiceProvider) Get() (*[]User, error) {
	var (
		user []User
	)

	o := orm.NewOrm()
	qs := o.QueryTable("user")
	_, err := qs.Filter("isactive", true).All(&user)
	return &user, err
}

func (this *UserServiceProvider) GetByID(id uint32) (*User, error) {
	var (
		user User
	)

	o := orm.NewOrm()
	qs := o.QueryTable("user")
	err := qs.Filter("id", id).One(&user)
	return &user, err
}

func (this *UserServiceProvider) Delete(id uint32) (int64, error) {
	user := User{ID:id}

	o := orm.NewOrm()
	user.Isactive = false
	return o.Update(&user, "isactive")
}
