package entity

type BiliUser struct {
	//gorm.Model
	Model
	ID     uint   `gorm:"comment:主键;primarykey"`
	Space  string `gorm:"comment:个人空间;type:varchar(255)"`
	Avatar string `gorm:"comment:头像;type:varchar(255)"`
	Name   string `gorm:"comment:用户名;type:varchar(255)"`
}
