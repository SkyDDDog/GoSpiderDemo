package Demo02

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"spider02/entity"
	"time"

	"strconv"
)

func InitDb() *gorm.DB {
	myLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)
	dsn := "root:9738faq@tcp(127.0.0.1:3306)/demo02?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: myLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true}},
	)
	return db
}

func IsMemberInDb(db *gorm.DB, uid uint) bool {
	var user entity.BiliUser
	result := db.First(&user, strconv.Itoa(int(uid)))

	if result.RowsAffected != 0 {
		return true
	} else {
		return false
	}
}
