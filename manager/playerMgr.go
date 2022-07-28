package manager

import "github.com/lqgl/tinywork/player"

// PlayerMgr 维护在线玩家
type PlayerMgr struct {
	players map[uint64]*player.Player
	addPCh  chan *player.Player
}

func NewPlayerMgr() *PlayerMgr {
	return &PlayerMgr{
		players: make(map[uint64]*player.Player),
		addPCh:  make(chan *player.Player, 1),
	}
}

func (pm *PlayerMgr) Add(p *player.Player) {
	if pm.players[p.UId] != nil {
		return
	}
	pm.players[p.UId] = p
	go p.Run()
}

// Del ...
func (pm *PlayerMgr) Del(p *player.Player) {
	delete(pm.players, p.UId)
}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <-pm.addPCh:
			pm.Add(p)
		}
	}
}

// GetPlayer ...
func (pm *PlayerMgr) GetPlayer(uId uint64) *player.Player {
	p, ok := pm.players[uId]
	if ok {
		return p
	}
	return nil
}
