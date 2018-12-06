package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Teacher struct {
	Id         int       `orm:"column(id);auto"`
	Name       string    `orm:"column(name);size(128)"`
	Signature  string    `orm:"column(signature);size(128);null"`
	Picture    string    `orm:"column(picture);size(128)"`
	Price      float32   `orm:"column(price);null"`
	IssuePrice float32   `orm:"column(issue_price);null"`
	Views      uint64    `orm:"column(views);null"`
	Tag        string    `orm:"column(tag);size(128);null"`
	Intro      string    `orm:"column(intro);size(512)"`
	State      int8      `orm:"column(state);null" description:"0：辞职,1：可用,2：休息"`
	CreateTime time.Time `orm:"column(create_time);type(datetime);auto_now_add"`
}

func (t *Teacher) TableName() string {
	return "teacher"
}

func init() {
	orm.RegisterModel(new(Teacher))
}

// AddTeacher insert a new Teacher into database and returns
// last inserted Id on success.
func AddTeacher(m *Teacher) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetTeacherById retrieves Teacher by Id. Returns error if
// Id doesn't exist
func GetTeacherById(id int) (v *Teacher, err error) {
	o := orm.NewOrm()
	v = &Teacher{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllTeacher retrieves all Teacher matches certain condition. Returns empty list if
// no records exist
func GetAllTeacher(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Teacher))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Teacher
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateTeacher updates Teacher by Id and returns error if
// the record to be updated doesn't exist
func UpdateTeacherById(m *Teacher) (err error) {
	o := orm.NewOrm()
	v := Teacher{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteTeacher deletes Teacher by Id and returns error if
// the record to be deleted doesn't exist
func DeleteTeacher(id int) (err error) {
	o := orm.NewOrm()
	v := Teacher{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Teacher{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
