package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type Document struct {
	Id            int       `orm:"column(id);auto"`
	UserId        int       `orm:"column(user_id)"`
	Avatar        string    `orm:"column(avatar);size(128);null"`
	Nickname      string    `orm:"column(nickname);size(128)"`
	Sex           uint      `orm:"column(sex);null" description:"0:男;1:女"`
	Birthday      time.Time `orm:"column(birthday);type(datetime);auto_now_add"`
	Height        int       `orm:"column(height);null"`
	Weight        int       `orm:"column(weight);null"`
	Job           string    `orm:"column(job);size(128);null"`
	Character     string    `orm:"column(character);size(128);null"`
	FavoriteStyle string    `orm:"column(favorite_style);size(256);null"`
	CreateTime    time.Time `orm:"column(create_time);type(datetime);auto_now_add"`
	UpdateTime    time.Time `orm:"column(update_time);type(datetime);auto_now_add"`
}

func (t *Document) TableName() string {
	return "document"
}

func init() {
	orm.RegisterModel(new(Document))
}

// AddDocument insert a new Document into database and returns
// last inserted Id on success.
func AddDocument(m *Document) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDocumentById retrieves Document by Id. Returns error if
// Id doesn't exist
func GetDocumentById(id int) (v *Document, err error) {
	o := orm.NewOrm()
	v = &Document{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDocument retrieves all Document matches certain condition. Returns empty list if
// no records exist
func GetAllDocument(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Document))
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

	var l []Document
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

// UpdateDocument updates Document by Id and returns error if
// the record to be updated doesn't exist
func UpdateDocumentById(m *Document) (err error) {
	o := orm.NewOrm()
	v := Document{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDocument deletes Document by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDocument(id int) (err error) {
	o := orm.NewOrm()
	v := Document{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Document{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
