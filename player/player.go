package player

type Player struct {
	UId        uint64
	FriendList []uint64 // 好友
	// HandlerParamCh chan HandlerParam
	handlers map[string]Handler // 相当于一个事件路由
}

func NewPlayer() *Player {
	p := &Player{
		UId:        0,
		FriendList: make([]uint64, 100),
		handlers:   make(map[string]Handler),
	}
	p.HandlerRegister()
	return p
}

func (p *Player) Run() {
	for {
		select {
		//case handlerParam := <-p.HandlerParamCh:
		//	if fn, ok := p.handlers[handlerParam]; ok {
		//		fn(handlerParam)
		//	}
		}
	}
}
