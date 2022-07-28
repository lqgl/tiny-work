package player

import (
	"github.com/golang/protobuf/proto"
	"github.com/lqgl/tasty"
	"github.com/lqgl/tinywork/logger"
	"github.com/lqgl/tinywork/network"
	player "github.com/lqgl/tinywork/network/protocol/gen/proto"
)

type Handler func(packet *network.Message)

// AddFriend 添加好友
func (p *Player) AddFriend(packet *network.Message) {
	req := &player.AddFriendReq{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		logger.Logger.DebugF("%v\n", err)
		return
	}
	if !tasty.CheckInSlice(req.UId, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UId)
	}
}

// DelFriend 删除好友
func (p *Player) DelFriend(packet *network.Message) {
	req := &player.DelFriendReq{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		logger.Logger.DebugF("%v\n", err)
		return
	}
	p.FriendList = tasty.DelOneInSlice(req.UId, p.FriendList)
}

// ResolveChatMsg 处理消息逻辑
func (p *Player) ResolveChatMsg(packet *network.Message) {
	req := &player.SendChatMsgReq{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		logger.Logger.DebugF("%v\n", err)
		return
	}
	chatMsg := req.Msg.Content
	logger.Logger.InfoF(chatMsg)
	//todo 收到消息 然后转发给客户端（当你的好友给你发消息的情况）
}
