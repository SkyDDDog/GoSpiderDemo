package entity

import "time"

type SubComment struct {
	Model
	ID       uint      `gorm:"comment:主键;primarykey"`
	Cid      uint      `gorm:"comment:主评论id;type:bigint"`
	Uid      uint      `gorm:"comment:用户uid;type:bigint"`
	Message  string    `gorm:"comment:评论内容"`
	PushTime time.Time `gorm:"comment:发布时间;type:datetime"`
	Likes    int       `gorm:"comment:点赞数;type:int"`
}
