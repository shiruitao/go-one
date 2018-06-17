/*
 * MIT License
 *
 * Copyright (c) 2018 Shi Ruitao.
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
 *     Initial: 2018/04/19        Shi Ruitao
 */

package models

import (
	"github.com/astaxie/beego/orm"
)

type WareServiceProvider struct{}

var WareService *WareServiceProvider

type Commodity struct {
	ID        uint32    `orm:"column(id);pk;auto" json:"id"`
	Name      string    `orm:"column(title)" json:"name"`
	Desc      string    `orm:"column(detail)" json:"desc"`
	Price     float32   `orm:"column(price)" json:"price"`
	Category  string    `orm:"column(class)" json:"category"`
	Avatar    string    `orm:"column(image)" json:"avatar"`
}

func init() {
	orm.RegisterModel(new(Commodity))
}

func (this *WareServiceProvider) AddWare(info *Commodity) (uint32, error) {
	var (
		ware Commodity
	)

	ware.Name = info.Name
	ware.Desc = info.Desc
	ware.Price = info.Price
	ware.Category = info.Category
	ware.Avatar = info.Avatar

	o := orm.NewOrm()
	id, err := o.Insert(&ware)
	return uint32(id), err
}

func (this *WareServiceProvider) GetAll() (*[]Commodity, int64, error) {
	var (
		ware []Commodity
	)
	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM commodity ORDER BY id DESC").QueryRows(&ware)
	return &ware, num, err
}

func (this *WareServiceProvider) Delete(id uint32) error {
	ware := Commodity{ID: id}

	o := orm.NewOrm()
	_, err := o.Delete(&ware)
	return err
}

func (this *WareServiceProvider) GetRecommend() (*[]Commodity, int64, error) {
	var (
		ware []Commodity
	)
	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM ware WHERE status = 2 ORDER BY id DESC").QueryRows(&ware)
	return &ware, num, err
}

func (this *WareServiceProvider) UpdateWare(info *Commodity) error {
	o := orm.NewOrm()

	ware := Commodity{ID: info.ID}
	ware.Name = info.Name
	ware.Desc = info.Desc
	ware.Price = info.Price
	ware.Category = info.Category
	ware.Avatar = info.Avatar

	_, err := o.Update(&ware, "name", "desc", "price", "saleprice", "inventory", "category", "status", "avatar", "image", "detailpic", "updated")
	return err
}

func (this *WareServiceProvider) StatusWare(id uint32, status int8) error {
	o := orm.NewOrm()

	ware := Commodity{ID: id}

	_, err := o.Update(&ware, "status")
	return err
}

func (*WareServiceProvider) GetByID(id []uint32) (Commodity, error) {
	var (
		ware Commodity
	)
	o := orm.NewOrm()
	_, err := o.QueryTable("ware").Filter("id__in", id).All(&ware)
	return ware, err
}
