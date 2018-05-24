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

func (this *UserServiceProvider) GetUserByAge() (int64, int64, int64, int64, int64, error) {
	var (
		user []User
	)
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	num1, err := qs.Filter("age__gte", 0).Filter("age__lte", 7).All(&user)
	num2, err := qs.Filter("age__gte", 7).Filter("age__lte", 17).All(&user)
	num3, err := qs.Filter("age__gte", 18).Filter("age__lte", 40).All(&user)
	num4, err := qs.Filter("age__gte", 41).Filter("age__lte", 65).All(&user)
	num5, err := qs.Filter("age__gte", 66).All(&user)
	return num1, num2, num3, num4, num5, err
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

func (this *UserServiceProvider) GetUserByArea(area string) (*[]User, int64, error) {
	var (
		user []User
	)
	o := orm.NewOrm()

	num, err := o.Raw("SELECT * FROM user WHERE area = ?", area).QueryRows(&user)
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