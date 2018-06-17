package models

import (
	"github.com/astaxie/beego/orm"
)

type QuestionServiceProvider struct{}

var QuestionService *QuestionServiceProvider

type Question struct {
	ID    uint32 `orm:"column(id);pk;auto"`
	One   int    `orm:"column(one)" json:"answer1"`
	Two   int    `orm:"column(two)" json:"answer2"`
	Three int    `orm:"column(three)" json:"answer3"`
	Four  int    `orm:"column(four)" json:"answer4"`
	Five  int    `orm:"column(five)" json:"answer5"`
	Six   int    `orm:"column(six)" json:"answer6"`
	Seven int    `orm:"column(seven)" json:"answer7"`
	Eight int    `orm:"column(eight)" json:"answer8"`
	Nine  int    `orm:"column(nine)" json:"answer9"`
	Ten   int    `orm:"column(ten)" json:"answer10"`
}

func init() {
	orm.RegisterModel(new(Question))
}

func (*QuestionServiceProvider) Add(q *Question) {
	qu := Question{
		One:   q.One,
		Two:   q.Two,
		Three: q.Three,
		Four:  q.Four,
		Five:  q.Five,
		Six:   q.Six,
		Seven: q.Seven,
		Eight: q.Eight,
		Nine:  q.Nine,
		Ten:   q.Ten,
	}

	o := orm.NewOrm()
	o.Insert(&qu)
}
