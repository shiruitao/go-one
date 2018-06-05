package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"github.com/shiruitao/go-one/application/hjjc/common"
)

type UserServiceProvider struct{}

var UserService *UserServiceProvider

type User struct {
	ID       uint32 `orm:"column(UserID);pk;auto"`
	Name     string `orm:"column(UserName)" json:"name"`
	RealName string `orm:"column(RealName)" json:"real_name"`
	Password string `orm:"column(Userpwd)" json:"password"`
	UserRole string `orm:"column(UserRole)"`
}

func init() {
	orm.RegisterModel(new(User))
}

func (this *UserServiceProvider) CreateUser(userInfo *User) (int64, error) {
	var (
		user User
	)
	user.Name = userInfo.Name
	user.RealName = userInfo.RealName
	user.Password = userInfo.Password
	user.UserRole = "普通用户"

	o := orm.NewOrm()

	return o.Insert(&user)
}

func (this *UserServiceProvider) Login(name, password string) (User, bool, error) {
	var (
		user User
	)

	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("UserName", name).One(&user)
	if err != nil {
		return user, false, err
	}
	if user.Password != password {
		return user, false, errors.New(common.ErrWrongPass)
	}

	return user, true, err
}

func (*UserServiceProvider) Delete(id uint32) error {
	user := User{ID: id}
	o := orm.NewOrm()
	_, err := o.Delete(&user)
	return err
}

func (*UserServiceProvider) Update(info *User) error {
	user := User{
		ID: info.ID,
		Name: info.Name,
		RealName: info.RealName,
		Password: info.Password,
	}

	o := orm.NewOrm()
	_, err := o.Update(&user, "UserName", "RealName", "Userpwd")

	return err
}

func (*UserServiceProvider) Get(id uint32) (*User, error) {
	user := User{ID:id}

	o := orm.NewOrm()

	err := o.Read(&user)
	return &user, err
}

func (*UserServiceProvider) GetAll() (*[]User, error) {
	var user []User
	o := orm.NewOrm()

	_, err := o.QueryTable("user").Exclude("id", 1).All(&user)
	return &user, err
}
