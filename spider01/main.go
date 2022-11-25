package main

import (
	"fmt"
	"spider01/entity"
	"spider01/service"
	"sync"
	"time"
)

func main() {
	//db := util.InitDb()
	//db.AutoMigrate(&entity.Fzu{})
	var wg sync.WaitGroup
	urlChan := make(chan string, 10)
	objChan := make(chan entity.Fzu, 1)
	resultChan := make(chan string, 1)
	start := time.Now()
	go service.GetFzuPages(urlChan)
	for i := 0; i < 10; i++ {
		go service.FzuSpider(&wg, urlChan, objChan)

	}
	go service.InsertSpider(&wg, objChan, resultChan)
	wg.Wait()
	for i := range resultChan {
		fmt.Println("数据", i, "已插入数据库")
	}
	fmt.Println("执行时间: ", time.Now().Sub(start))
	fmt.Println("done")
}
