package entity

import "time"

type BiliComment struct {
	Model
	ID       uint      `gorm:"comment:主键;primarykey"`
	Uid      uint      `gorm:"comment:用户uid;type:bigint"`
	Message  string    `gorm:"comment:评论内容"`
	PushTime time.Time `gorm:"comment:发布时间;type:datetime"`
	Likes    int       `gorm:"comment:点赞数;type:int"`
}
