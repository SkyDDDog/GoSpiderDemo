package main

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"spider02/entity"
	"spider02/entity/dto"
	"spider02/service"
	util "spider02/util"
	strconv "strconv"
	"time"
)

func main() {
	db := util.InitDb()
	db.AutoMigrate(&entity.BiliUser{})
	start := time.Now()
	// 每个next 20条数据
	url := "https://api.bilibili.com/x/v2/reply/main?type=1&oid=21071819&next=1"
	jsonStr := service.GetJson(url)
	for i := 2; !gjson.Get(jsonStr, "data.cursor.is_end").Bool(); i++ {
		basePath := "data.replies."
		subPath := ".replies."
		contentPath := ".content.message"
		var repliesJson string
		var subRepliesJson string
		var content string
		var subContent string
		var replies dto.Replies
		var subReplies dto.Replies
		for i := 0; i < 20; i++ {
			index := strconv.Itoa(i)
			repliesJson = gjson.Get(jsonStr, basePath+index).String()
			if repliesJson == "" {
				break
			}
			content = gjson.Get(jsonStr, basePath+index+contentPath).String()
			json.Unmarshal([]byte(repliesJson), &replies)
			memberModel := dto.MemberDTO2Model(replies.Member)
			replyModel := dto.ReplyDTO2Model(replies, content)
			db.Create(&replyModel)
			db.Order("update_date desc").First(&memberModel)
			//fmt.Println(memberModel.ID)
			//user, _ := db.Get(strconv.Itoa(int(memberModel.ID)))
			//fmt.Println(user)
			if !util.IsMemberInDb(db, memberModel.ID) {
				db.Create(&memberModel)
			}
			//fmt.Println(memberModel)
			//fmt.Println(replyModel)
			for j := 0; j < 20; j++ {
				subIndex := strconv.Itoa(j)
				subRepliesJson = gjson.Get(jsonStr, basePath+index+subPath+subIndex).String()
				if subRepliesJson == "" {
					break
				}

				subContent = gjson.Get(jsonStr, basePath+index+subPath+subIndex+contentPath).String()
				json.Unmarshal([]byte(subRepliesJson), &subReplies)
				db.Last(&replyModel)
				subReplyModel := dto.SubReplyDTO2Model(replyModel.ID, subReplies, subContent)
				memberModel = dto.MemberDTO2Model(subReplies.Member)
				db.Order("update_date desc").Limit(1).Find(&memberModel)
				//fmt.Println(memberModel.ID)
				if !util.IsMemberInDb(db, memberModel.ID) {
					db.Create(&memberModel)
				}
				db.Create(&subReplyModel)
				//fmt.Println(subReplyModel)
			}

		}
		// 更新请求json
		url = "https://api.bilibili.com/x/v2/reply/main?type=1&oid=21071819&next=" + strconv.Itoa(i)
		jsonStr = service.GetJson(url)
	}

	fmt.Println("执行时间: ", time.Now().Sub(start))
	fmt.Println("done")
}
