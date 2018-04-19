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
	ID       uint32 `orm:"column(id);pk;auto" json:"id"`
	Name     string `orm:"column(name)" json:"name"`
	Desc     string `orm:"column(desc)" json:"desc"`
	Price    float32 `orm:"column(price)" json:"price"`
	SalePrice float32 `orm:"column(saleprice)" json:"sale_price"`
	Inventory int32 `orm:"column(inventory)" json:"inventory"`
	CategoryID uint32 `orm:"column(category)" json:"category_id"`
	Avatar  string `orm:"column(avatar)" json:"avatar"`
	Picture string `orm:"column(picture)" json:"picture"`
	DetailPic string `orm:"column(detailpic)" json:"detail_pic"`
	Status   string `orm:"column(status)" json:"status"`
	Created  time.Time `orm:"column(created)" json:"created"`
	Updated  time.Time `orm:"column(updated)" json:"updated"`
}

func (this *WareServiceProvider) init() {
	o := orm.NewOrm()
	o.Using("store")
}

func (this *WareServiceProvider) AddWare(info *Ware) (uint32, error){
	o := orm.NewOrm()
	id, err :=o.Insert(&info)
	return uint32(id), err
}
