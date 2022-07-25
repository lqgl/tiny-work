package main

import "github.com/lqgl/tinywork/network"

func main() {
	server := network.NewServer(":8023", "tcp6")
	server.Run()
	select {} // 用于阻塞 main 函数
}
