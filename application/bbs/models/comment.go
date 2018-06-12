package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type CommentServiceProvider struct{}

var CommentService *CommentServiceProvider

type Comment struct {
	ID        uint32    `orm:"column(id);pk;auto"`
	ArtID     uint32    `orm:"column(artid)" json:"art_id"`
	Content   string    `orm:"column(content)" json:"content"`
	CreatorID uint32    `orm:"column(creatorid)" json:"creator_id"`
	Creator   string    `orm:"column(creator)" json:"creator"`
	RepliedID uint32    `orm:"column(repliedid)" json:"replied_id"`
	Replied   string    `orm:"column(replied)" json:"replied"`
	RepContent   string    `orm:"column(repcontent)" json:"rep_content"`
	IsActive  bool      `orm:"column(isactive)"`
	File      string    `orm:"column(file)" json:"file"`
	RepFile      string    `orm:"column(repfile)" json:"rep_file"`
	Created   time.Time `orm:"column(created)"`
}

func (*CommentServiceProvider) Add(info *Comment) error {
	var (
		c Comment
	)
	c.ArtID = info.ArtID
	c.Content = info.Content
	c.CreatorID = info.CreatorID
	c.Creator = info.Creator
	c.RepliedID = info.RepliedID
	c.Replied = info.Replied
	c.RepContent = info.RepContent
	c.RepFile = info.RepFile
	c.IsActive = true
	c.File = info.File

	o := orm.NewOrm()
	_, err := o.Insert(&c)
	return err
}

func (*CommentServiceProvider) Get(id uint32) (*Comment, error) {
	var comment Comment
	o := orm.NewOrm()
	err := o.QueryTable("comment").Filter("artid", id).One(&comment)

	return &comment, err
}

func (*CommentServiceProvider) Delete(id uint32) (int64, error) {
	c := Comment{ID:id}
	return orm.NewOrm().Delete(&c)
}
