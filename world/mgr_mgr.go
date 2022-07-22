package world

import "github.com/lqgl/tinywork/manager"

type MgrMgr struct {
	Pm manager.PlayerMgr
}

func NewMgrMgr() *MgrMgr {
	m := &MgrMgr{
		Pm: manager.PlayerMgr{},
	}
	return m
}

var MM *MgrMgr
