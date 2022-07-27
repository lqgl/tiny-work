package world

import messageId "github.com/lqgl/tinywork/network/protocol/gen/proto"

func (mm *MgrMgr) HandlerRegister() {
	mm.Handlers[messageId.MessageId_CreatePlayerReq] = mm.CreatePlayer
	mm.Handlers[messageId.MessageId_LoginReq] = mm.UserLogin
}
