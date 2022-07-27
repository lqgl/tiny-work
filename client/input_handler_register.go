package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/lqgl/tinywork/network"
	messageId "github.com/lqgl/tinywork/network/protocol/gen/proto"
)

func (c *Client) InputHandlerRegister() {
	c.inputHandlers[messageId.MessageId_CreatePlayerReq.String()] = c.CreatePlayerReq
	c.inputHandlers[messageId.MessageId_LoginReq.String()] = c.Login
	c.inputHandlers[messageId.MessageId_AddFriendReq.String()] = c.AddFriendReq
	c.inputHandlers[messageId.MessageId_DelFriendReq.String()] = c.DelFriendReq
	c.inputHandlers[messageId.MessageId_SendChatMsgReq.String()] = c.SendChatMsgReq
}

func (c *Client) GetMessageIdByCmd(cmd string) messageId.MessageId {
	mid, ok := messageId.MessageId_value[cmd]
	if ok {
		return messageId.MessageId(mid)
	}
	return messageId.MessageId_None
}

func (c *Client) Transport(id messageId.MessageId, message proto.Message) {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return
	}
	c.cli.ChMsg <- &network.Message{
		Id:   uint64(id),
		Data: bytes,
	}
}
