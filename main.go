package main

import "github.com/lqgl/tinywork/world"

func main() {
	world.MM = world.NewMgrMgr()
	world.MM.Pm.Run()
}
