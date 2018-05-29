package models

import (
	"errors"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/shiruitao/go-one/application/graduation/common"
	"github.com/shiruitao/go-one/application/graduation/utility"
)

type UserServiceProvider struct{}

var UserService *UserServiceProvider

type (
	User struct {
		ID         uint32    `orm:"column(id);pk;auto"`
		Name       string    `orm:"column(name)" json:"name"`
		Sex        string    `orm:"column(sex)" json:"sex"`
		Number     string    `orm:"column(number)" json:"number"`
		Password   string    `orm:"column(password)";type:varchar(128)" json:"password"`
		Profession string    `orm:"column(profession)" json:"profession"`
		Department string    `orm:"column(department)" json:"department"`
		Major      string    `orm:"column(major)" json:"major"`
		Class      string    `orm:"column(class)" json:"class"`
		Year       int32      `orm:"column(year)" json:"year"`
		Power      int8      `orm:"column(power)"`
		LastLogin  time.Time `orm:"column(lastlogin);type(datetime)"`
		Created    time.Time `orm:"column(created);auto_now_add;type(datetime)"`
	}
)

//func init() {
//	orm.RegisterModel(new(User))
//}

func (this *UserServiceProvider) CreateUser(info *User) (int64, error) {
	var (
		user User
		id   int64
		err  error
	)
	user.Name = info.Name
	user.Sex = info.Sex
	user.Number = info.Number
	user.Profession = info.Profession
	user.Department = info.Department
	user.Major = info.Major
	user.Class = info.Class
	user.Year = info.Year
	user.Power = info.Power
	user.LastLogin = time.Now()

	if len(info.Password) < 6 || len(info.Password) > 30 {
		err = errors.New(common.ErrWrongPass)
		return id, err
	}

	hash, err := utility.GenerateHash(info.Password)
	user.Password = string(hash)
	if err != nil {
		return id, err
	}
	o := orm.NewOrm()

	return o.Insert(&user)
}

func (this *UserServiceProvider) ChangeUserInfo(id uint32, info *User) (int64, error) {
	var (
		userInfo User
	)

	userInfo.ID = info.ID
	userInfo.Name = info.Name
	userInfo.Sex = info.Sex
	userInfo.Number = info.Number
	userInfo.Profession = info.Profession
	userInfo.Department = info.Department
	userInfo.Major = info.Major
	userInfo.Class = info.Class
	userInfo.Year = info.Year

	o := orm.NewOrm()

	return o.Update(&userInfo)
}

func (this *UserServiceProvider) Login(name, password string) (User, bool, error) {
	var (
		user User
	)

	o := orm.NewOrm()
	err := o.QueryTable("user").Filter("Number", name).One(&user)
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
