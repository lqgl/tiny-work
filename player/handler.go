package player

import (
	"fmt"
	"github.com/lqgl/tinywork/chat"
	"github.com/lqgl/tinywork/function"
)

type Handler func(interface{})

// AddFriend 添加好友
func (p *Player) AddFriend(data interface{}) {
	fId := data.(uint64)
	if !function.CheckInNumberSlice(fId, p.FriendList) {
		p.FriendList = append(p.FriendList, fId)
	}
}

// DelFriend 删除好友
func (p *Player) DelFriend(data interface{}) {
	fId := data.(uint64)
	p.FriendList = function.DelEleInSlice(fId, p.FriendList)
}

// ResolveChatMsg 处理消息逻辑
func (p *Player) ResolveChatMsg(data interface{}) {
	chatMsg := data.(chat.Msg)
	fmt.Println(chatMsg)
	//todo 收到消息 然后转发给客户端（当你的好友给你发消息的情况）
}
