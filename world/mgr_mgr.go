package world

import (
	"github.com/lqgl/tinywork/manager"
	"github.com/lqgl/tinywork/network"
	messageId "github.com/lqgl/tinywork/network/protocol/gen/proto"
)

type MgrMgr struct {
	Pm              *manager.PlayerMgr
	Server          *network.Server
	Handlers        map[messageId.MessageId]func(message *network.SessionPacket)
	sessionPacketCh chan *network.SessionPacket
}

func NewMgrMgr() *MgrMgr {
	mm := &MgrMgr{
		Pm: manager.NewPlayerMgr(),
	}
	mm.Server = network.NewServer(":8023", "tcp6")
	mm.Server.OnSessionPacket = mm.OnSessionPacket // 初始化 server.OnSessionPacket
	mm.Handlers = make(map[messageId.MessageId]func(message *network.SessionPacket))
	return mm
}

var MM *MgrMgr

func (mm *MgrMgr) Run() {
	mm.HandlerRegister() // 注册处理消息的路由
	go mm.Server.Run()
	go mm.Pm.Run()
}

// OnSessionPacket 处理客户端发送的消息（服务器接收到的）
func (mm *MgrMgr) OnSessionPacket(packet *network.SessionPacket) {
	if handler, ok := mm.Handlers[messageId.MessageId(packet.Msg.Id)]; ok {
		handler(packet)
		return
	}
	if p := mm.Pm.GetPlayer(packet.Sess.UId); p != nil {
		p.HandlerParamCh <- packet.Msg
	}
}
