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
	"time"

	"github.com/astaxie/beego/orm"
)

type WareServiceProvider struct{}

var WareService *WareServiceProvider

type Ware struct {
	ID        uint32    `orm:"column(id);pk;auto" json:"id"`
	Name      string    `orm:"column(name)" json:"name"`
	Desc      string    `orm:"column(desc)" json:"desc"`
	Price     float32   `orm:"column(price)" json:"price"`
	SalePrice float32   `orm:"column(saleprice)" json:"sale_price"`
	Inventory int32     `orm:"column(inventory)" json:"inventory"`
	Category  string    `orm:"column(category)" json:"category"`
	Avatar    string    `orm:"column(avatar)" json:"avatar"`
	Image     string    `orm:"column(image)" json:"image"`
	DetailPic string    `orm:"column(detailpic)" json:"detail_pic"`
	Status    int8      `orm:"column(status)" json:"status"` // 0 -> hide; 1 -> normal; 2 -> recommend
	Updated   time.Time `orm:"column(updated)" json:"updated"`
	Created   time.Time `orm:"column(created);auto_now_add;type(datetime)" json:"created"`
}

func init() {
	orm.RegisterModel(new(Ware))
}

func (this *WareServiceProvider) AddWare(info *Ware) (uint32, error) {
	var (
		ware Ware
	)

	ware.Name = info.Name
	ware.Desc = info.Desc
	ware.Price = info.Price
	ware.SalePrice = info.SalePrice
	ware.Inventory = info.Inventory
	ware.Category = info.Category
	ware.Status = info.Status
	ware.Avatar = info.Avatar
	ware.Image = info.Image
	ware.DetailPic = info.DetailPic

	o := orm.NewOrm()
	id, err := o.Insert(&ware)
	return uint32(id), err
}

func (this *WareServiceProvider) GetAll() (*[]Ware, int64, error) {
	var (
		ware []Ware
	)
	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM ware WHERE status IN (1,2) ORDER BY id DESC").QueryRows(&ware)
	return &ware, num, err
}

func (this *WareServiceProvider) GetRecommend() (*[]Ware, int64, error) {
	var (
		ware []Ware
	)
	o := orm.NewOrm()
	num, err := o.Raw("SELECT * FROM ware WHERE status = 2 ORDER BY id DESC").QueryRows(&ware)
	return &ware, num, err
}

func (this *WareServiceProvider) UpdateWare(info *Ware) error {
	o := orm.NewOrm()

	ware := Ware{ID: info.ID}
	ware.Name = info.Name
	ware.Desc = info.Desc
	ware.Price = info.Price
	ware.SalePrice = info.SalePrice
	ware.Inventory = info.Inventory
	ware.Category = info.Category
	ware.Status = info.Status
	ware.Avatar = info.Avatar
	ware.Image = info.Image
	ware.DetailPic = info.DetailPic
	ware.Updated = time.Now()

	_, err := o.Update(&ware, "name", "desc", "price", "saleprice", "inventory", "category", "status", "avatar", "image", "detailpic", "updated")
	return err
}

func (this *WareServiceProvider) StatusWare(id uint32, status int8) error {
	o := orm.NewOrm()

	ware := Ware{ID: id}
	ware.Status = status

	_, err := o.Update(&ware, "status")
	return err
}
