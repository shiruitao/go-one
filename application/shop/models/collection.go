/*
 * MIT License
 *
 * Copyright (c) 2018 SmartestEE Co., Ltd..
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2018/05/30        Shi Ruitao
 */

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type CollectionServiceProvider struct{}

var CollectionService CollectionServiceProvider

type Collection struct {
	ID       uint32    `orm:"column(id);pk;auto"`
	UserID   uint32    `orm:"column(userid)"`
	WareID   uint64    `orm:"column(wareid)" json:"ware_id"`
	Created  time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Collection))
}

func (*CollectionServiceProvider) Add(info *Collection) (int64, error) {
	var (
		c Collection
	)

	c.UserID = info.UserID
	c.WareID = info.WareID
	o := orm.NewOrm()

	return o.Insert(&c)
}

func (*CollectionServiceProvider) Delete(id uint32) (int64, error) {
	c := Collection{ID: id}
	return orm.NewOrm().Delete(&c)
}

func (*CollectionServiceProvider) Get(userID uint32) (*[]Collection, error) {
	var c []Collection

	o := orm.NewOrm()
	qs := o.QueryTable("collection")
	_, err := qs.Filter("userid", userID).All(&c)
	return &c, err
}
