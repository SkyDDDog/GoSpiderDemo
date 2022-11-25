package entity

import (
	"time"
)

type Fzu struct {
	//gorm.Model
	Model
	Author       string    `gorm:"comment:作者;type:varchar(255)"`
	Title        string    `gorm:"comment:标题;type:varchar(255)"`
	Read_num     int       `gorm:"comment:阅读量;type:int"`
	Content      string    `gorm:"comment:正文;type:longtext"`
	Publish_time time.Time `gorm:"comment:发布时间;type:date"`
	Link         string    `gorm:"comment:链接;type:varchar(255)"`
}
