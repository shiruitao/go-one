package models

import (
	"errors"
)

const (
	// Database name.
	Database = "bbs"

	// Increase add 1.
	Increase = 1
	// Decrease submit 1
	Decrease = -1

	// CommentDeleted represents the comment which is deleted.
	CommentDeleted = -1
	// CommentUnread represents the comment which is not read.
	CommentUnread = 0
	// CommentRead represents the comment which id read.
	CommentRead = 1
)

// InvalidObjectId
var (
	InvalidObjectId = errors.New("invalid input to ObjectIdHex: ")
)
