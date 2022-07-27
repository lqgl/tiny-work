package main

import "github.com/lqgl/tinywork/world"

func main() {
	world.MM = world.NewMgrMgr()
	go world.MM.Run()
	select {}
}
