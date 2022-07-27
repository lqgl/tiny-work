package player

import (
	"github.com/lqgl/tinywork/network"
	"github.com/lqgl/tinywork/network/protocol/gen/proto"
)

type Player struct {
	UId            uint64
	FriendList     []uint64 // 好友
	HandlerParamCh chan *network.Message
	handlers       map[proto.MessageId]Handler // 相当于一个事件路由
	Session        *network.Session
}

func NewPlayer() *Player {
	p := &Player{
		UId:        0,
		FriendList: make([]uint64, 100),
		handlers:   make(map[proto.MessageId]Handler),
	}
	p.HandlerRegister()
	return p
}

func (p *Player) Run() {
	for {
		select {
		case handlerParam := <-p.HandlerParamCh:
			if fn, ok := p.handlers[proto.MessageId(handlerParam.Id)]; ok {
				fn(handlerParam)
			}
		}
	}
}
