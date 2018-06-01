package models

import (
	"time"
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
	IsActive  bool      `orm:"column(isactive)"`
	File      string    `orm:"column(file)" json:"file"`
	Created   time.Time `orm:"column(created)"`
}


