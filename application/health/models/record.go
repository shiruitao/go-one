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
 *     Initial: 2018/06/03        Shi Ruitao
 */

package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type RecordServiceProvider struct{}

var RecordService *RecordServiceProvider

type Record struct {
	ID      uint32    `orm:"column(id);pk;auto"`
	Name    string    `orm:"column(name)" json:"name"`
	Score   int       `orm:"column(score)" json:"score"`
	Result  string    `orm:"column(result)" json:"result"`
	Created time.Time `orm:"column(created);auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Record))
}

func (*RecordServiceProvider) Add(info *Record) (int64, error) {
	var record Record

	record.Name = info.Name
	record.Score = info.Score
	record.Result = info.Result
	o := orm.NewOrm()
	return o.Insert(&record)
}

func (*RecordServiceProvider) Get(name string) *[]Record {
	var record []Record
	o := orm.NewOrm()
	o.QueryTable("record").Filter("name", name).All(&record)
	return &record
}
