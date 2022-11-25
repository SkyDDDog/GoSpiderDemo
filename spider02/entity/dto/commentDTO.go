package dto

import (
	"spider02/entity"
	"strconv"
	"time"
)

type Replies struct {
	Oid     string
	Member  Member
	Content string
	ctime   string
	like    int
}

type Member struct {
	Mid    string
	Uname  string
	Avatar string
}

func MemberDTO2Model(member Member) entity.BiliUser {
	uid, _ := strconv.Atoi(member.Mid)
	var user = entity.BiliUser{
		ID:     uint(uid),
		Space:  "https://space.bilibili.com/" + member.Mid,
		Avatar: member.Avatar,
		Name:   member.Uname,
	}
	return user
}

func ReplyDTO2Model(replies Replies, content string) entity.BiliComment {
	timeStamp, _ := strconv.Atoi(replies.ctime)
	uid, _ := strconv.Atoi(replies.Member.Mid)
	unix := time.Unix(int64(timeStamp), 0)
	var comment = entity.BiliComment{
		Uid:      uint(uid),
		Message:  content,
		PushTime: unix,
		Likes:    replies.like,
	}
	return comment
}

func SubReplyDTO2Model(cid uint, replies Replies, content string) entity.SubComment {
	timeStamp, _ := strconv.Atoi(replies.ctime)
	uid, _ := strconv.Atoi(replies.Member.Mid)
	unix := time.Unix(int64(timeStamp), 0)
	var comment = entity.SubComment{
		Cid:      cid,
		Uid:      uint(uid),
		Message:  content,
		PushTime: unix,
		Likes:    replies.like,
	}
	return comment
}
