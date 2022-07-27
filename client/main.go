package main

func main() {
	c := NewClient()
	c.InputHandlerRegister()   // 注册输入处理路由
	c.MessageHandlerRegister() // 注册消息处理路由
	c.Run()
	select {}
}
