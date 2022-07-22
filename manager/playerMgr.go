package manager

import "github.com/lqgl/tinywork/player"

// PlayerMgr 维护在线玩家
type PlayerMgr struct {
	players map[uint64]player.Player
	addPCh  chan player.Player
}

func (pm *PlayerMgr) Add(p player.Player) {
	pm.players[p.UId] = p
}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <-pm.addPCh:
			pm.Add(p)
		}
	}
}
