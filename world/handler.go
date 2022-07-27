package world

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/lqgl/tinywork/network"
	player "github.com/lqgl/tinywork/network/protocol/gen/proto"
	logicPlayer "github.com/lqgl/tinywork/player"
)

func (mm *MgrMgr) CreatePlayer(message *network.SessionPacket) {
	msg := &player.CreateUserReq{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	fmt.Println("[MgrMgr.CreatePlayer]", msg)
	mm.SendMsg(uint64(player.MessageId_CreatePlayerResp), &player.CreateUserResp{}, message.Sess)
}

func (mm *MgrMgr) UserLogin(message *network.SessionPacket) {
	msg := &player.LoginReq{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	newPlayer := logicPlayer.NewPlayer()
	newPlayer.UId = 111
	newPlayer.HandlerParamCh = message.Sess.ChWrite
	message.Sess.IsPlayerOnline = true
	message.Sess.UId = newPlayer.UId
	newPlayer.Session = message.Sess
	mm.Pm.Add(newPlayer)
}

func (mm MgrMgr) SendMsg(id uint64, message proto.Message, session *network.Session) {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return
	}
	resp := &network.Message{
		Id:   id,
		Data: bytes,
	}
	session.SendMsg(resp)
}
