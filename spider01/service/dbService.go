package service

import (
	"spider01/entity"
	util "spider01/util"
	"sync"
)

func InsertSpider(wg *sync.WaitGroup, objChan chan entity.Fzu, result chan<- string) {
	db := util.InitDb()
	db.AutoMigrate(&entity.Fzu{})
	for fdyw := range objChan {
		db.Create(&fdyw)
		result <- fdyw.Link
		wg.Done()
	}
	close(result)

}
