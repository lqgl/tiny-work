package main

import messageId "github.com/lqgl/tinywork/network/protocol/gen/proto"

func (c *Client) MessageHandlerRegister() {
	c.messageHandlers[messageId.MessageId_CreatePlayerResp] = c.OnCreatePlayerResp
	c.messageHandlers[messageId.MessageId_LoginResp] = c.OnLoginRsp
	c.messageHandlers[messageId.MessageId_AddFriendResp] = c.OnAddFriendResp
	c.messageHandlers[messageId.MessageId_DelFriendResp] = c.OnDelFriendResp
	c.messageHandlers[messageId.MessageId_SendChatMsgResp] = c.OnSendChatMsgRsp

}
