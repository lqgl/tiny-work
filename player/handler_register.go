package player

import messageId "github.com/lqgl/tinywork/network/protocol/gen/proto"

// HandlerRegister 注册事件
func (p *Player) HandlerRegister() {
	p.handlers[messageId.MessageId_AddFriendReq] = p.AddFriend
	p.handlers[messageId.MessageId_DelFriendReq] = p.DelFriend
	p.handlers[messageId.MessageId_SendChatMsgReq] = p.ResolveChatMsg
}
