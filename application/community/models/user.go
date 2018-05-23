package models

import (
	"time"

	"github.com/astaxie/beego/orm"

)

type UserServiceProvider struct{}

var UserService *UserServiceProvider

type (
	User struct {
		ID        uint32    `orm:"column(id);pk;auto"`
		Name      string    `orm:"column(name)" json:"name"`
		Sex       string    `orm:"column(sex)" json:"sex"`
		Age       int8      `orm:"column(age)" json:"age"`
		NumberID  string    `orm:"column(numberid)" json:"IDnumber"`
		Area      string    `orm:"column(area)" json:"area"`
		Address   string    `orm:"column(address)" json:"address"`
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
	user.Area = userInfo.Area
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

func (this *UserServiceProvider) GetUser() (*[]User, int64, error) {
	var (
		user []User
	)
	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM user ORDER BY id DESC").QueryRows(&user)
	return &user, num, err
}

func (this *UserServiceProvider) GetUserByAge(low, high int8) (*[]User, int64, error) {
	var (
		user []User
	)
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	num, err := qs.Filter("age__gte", low).Filter("age__lte", high).All(&user)
	return &user, num, err
}

func (this *UserServiceProvider) GetUserByName(name string) (*[]User, int64, error) {
	var (
		user []User
	)
	o := orm.NewOrm()

	name = "%" + name + "%"
	num, err := o.Raw("SELECT * FROM user WHERE name LIKE ?", name).QueryRows(&user)
	return &user, num, err
}

func (this *UserServiceProvider) GetUserByArea(name string) (*[]User, int64, error) {
	var (
		user []User
	)
	o := orm.NewOrm()

	num, err := o.Raw("SELECT * FROM user WHERE area = ?", name).QueryRows(&user)
	return &user, num, err
}

func (this *UserServiceProvider) ChangeInfo(userInfo *User) error {
	var (
		user User
	)
	user.ID = userInfo.ID
	user.Name = userInfo.Name
	user.Sex = userInfo.Sex
	user.Age = userInfo.Age
	user.NumberID = userInfo.NumberID
	user.Address = userInfo.Address

	o := orm.NewOrm()
	_, err := o.Update(&user, "name", "sex", "age", "numberid", "address")
	return err
}