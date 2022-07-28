package main

import (
	"github.com/golang/protobuf/proto"
	"github.com/lqgl/tinywork/logger"
	"github.com/lqgl/tinywork/network"
	player "github.com/lqgl/tinywork/network/protocol/gen/proto"
	"strconv"
)

type MessageHandler func(packet *network.ClientPacket)
type InputHandler func(param *InputParam)

func (c *Client) Login(param *InputParam) {
	logger.Logger.DebugF("Login input Handler print")
	logger.Logger.InfoF(param.Command)
	logger.Logger.InfoF("%v\n", param.Param)
	id := c.GetMessageIdByCmd(param.Command)
	if len(param.Param) != 2 {
		return
	}
	msg := &player.CreateUserReq{
		UserName: param.Param[0],
		Password: param.Param[1],
	}
	c.Transport(id, msg) // 将 msg 发送给 cli
}

// CreatePlayerReq 创建角色
func (c *Client) CreatePlayerReq(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)
	if len(param.Param) != 2 {
		return
	}
	msg := &player.CreateUserReq{
		UserName: param.Param[0],
		Password: param.Param[1],
	}
	c.Transport(id, msg) // 将 msg 发送给 cli
}

// OnCreatePlayerResp 创建角色结果响应
func (c *Client) OnCreatePlayerResp(packet *network.ClientPacket) {
	logger.Logger.InfoF("恭喜你创建角色成功")
}

// LoginReq 玩家登录请求
func (c *Client) LoginReq(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 2 {
		return
	}

	msg := &player.LoginReq{
		UserName: param.Param[0],
		Password: param.Param[1],
	}

	c.Transport(id, msg)
}

// OnLoginRsp 玩家登录结果响应
func (c *Client) OnLoginRsp(packet *network.ClientPacket) {
	rsp := &player.LoginResp{}

	err := proto.Unmarshal(packet.Msg.Data, rsp)
	if err != nil {
		return
	}

	logger.Logger.InfoF("登陆成功")
}

// AddFriendReq 添加好友请求
func (c *Client) AddFriendReq(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)
	if len(param.Param) != 1 || len(param.Param[0]) == 0 { //""
		return
	}
	parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}
	msg := &player.AddFriendReq{
		UId: parseUint,
	}
	c.Transport(id, msg)
}

// OnAddFriendResp 添加好友结果响应
func (c *Client) OnAddFriendResp(packet *network.ClientPacket) {
	req := &player.AddFriendResp{}
	err := proto.Unmarshal(packet.Msg.Data, req)
	if err != nil {
		logger.Logger.DebugF("%v\n", err)
		return
	}
	logger.Logger.InfoF("添加好友成功")
}

// DelFriendReq 删除好友请求
func (c *Client) DelFriendReq(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)
	if len(param.Param) != 1 || len(param.Param[0]) == 0 { //""
		return
	}
	parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}
	msg := &player.DelFriendReq{
		UId: parseUint,
	}
	c.Transport(id, msg)
}

// OnDelFriendResp 删除好友结果响应
func (c *Client) OnDelFriendResp(packet *network.ClientPacket) {
	req := &player.DelFriendResp{}
	err := proto.Unmarshal(packet.Msg.Data, req)
	if err != nil {
		logger.Logger.DebugF("%v\n", err)
		return
	}
	logger.Logger.InfoF("删除好友成功")
}

func (c *Client) SendChatMsgReq(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)
	if len(param.Param) != 1 || len(param.Param[0]) == 0 { //""
		return
	}
	parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	if err != nil {
		return
	}
	parseInt32, err := strconv.ParseInt(param.Param[2], 10, 32)
	if err != nil {
		return
	}

	msg := &player.SendChatMsgReq{
		UId: parseUint,
		Msg: &player.ChatMessage{
			Content: param.Param[1],
			Extra:   nil,
		},
		Category: int32(parseInt32),
	}

	c.Transport(id, msg)
}

func (c *Client) OnSendChatMsgRsp(packet *network.ClientPacket) {
	logger.Logger.InfoF("send  chat message success")

}
