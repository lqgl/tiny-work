package main

import "github.com/lqgl/tinywork/network"

func main() {
	client := network.NewClient(":8023", "tcp6")
	client.Run()
	select {} // 用于阻塞 main 函数
}
