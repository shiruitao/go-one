package models

import (
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/shiruitao/go-one/application/community/utility"
)

type UserServiceProvider struct{}

var UserService *UserServiceProvider

type (
	User struct {
		ID        uint32    `orm:"column(id);pk;auto"`
		Name      string    `orm:"column(name);not null;utf8_bin" json:"name"`
		Sex       string    `orm:"column(sex);not null;utf8_bin" json:"sex"`
		Age       int8      `orm:"column(age);not null;utf8_bin" json:"age"`
		NumberID  int64     `orm:"column(numberid);not null;utf8_bin" json:"IDnumber"`
		Address   string    `orm:"column(address);not null;utf8_bin" json:"address"`
		Created   time.Time `orm:"column(created);auto_now_add;type(datetime)"`
	}
)

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
	user.Sex = userInfo.Sex
	user.Age = userInfo.Age
	user.NumberID = userInfo.NumberID
	user.Address = userInfo.Address
	user.Created = time.Now()

	o := orm.NewOrm()

	id, err = o.Insert(&user)
	if err != nil {
		return id, err
	}

	return id, nil
}

func (this *UserServiceProvider) ChangeInfo(info *User) error {
	var (
		user User
	)

	o := orm.NewOrm()

	user = info
	_, err := o.Update(&user, "password")
	return err
}