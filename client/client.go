package main

import (
	"fmt"
	"github.com/lqgl/tinywork/network"
	messageId "github.com/lqgl/tinywork/network/protocol/gen/proto"
	"os"
	"syscall"
)

type Client struct {
	cli             *network.Client
	inputHandlers   map[string]InputHandler
	messageHandlers map[messageId.MessageId]MessageHandler
	console         *Console
	chInput         chan *InputParam
}

func NewClient() *Client {
	c := &Client{
		cli:             network.NewClient(":8023", "tcp6"),
		inputHandlers:   map[string]InputHandler{},
		messageHandlers: map[messageId.MessageId]MessageHandler{},
		console:         NewConsole(),
	}
	c.cli.OnMessage = c.OnMessage
	c.cli.ChMsg = make(chan *network.Message, 1)
	c.chInput = make(chan *InputParam, 1)
	c.console.chInput = c.chInput
	return c
}

func (c *Client) Run() {
	go func() {
		for {
			select {
			case input := <-c.chInput: // 接收到处理后的终端Input数据
				inputHandler := c.inputHandlers[input.Command]
				if inputHandler != nil {
					inputHandler(input)
				}
			}
		}
	}()
	go c.console.Run()
	go c.cli.Run()
}

// OnMessage 处理服务器返回的消息
func (c *Client) OnMessage(packet *network.ClientPacket) {
	if handler, ok := c.messageHandlers[messageId.MessageId(packet.Msg.Id)]; ok && packet.Msg != nil {
		handler(packet)
	}
}

func (c *Client) OnSystemSignal(signal os.Signal) bool {
	fmt.Printf("[Client] 收到信号 %v \n", signal)
	tag := true
	switch signal {
	case syscall.SIGHUP:
	// todo
	case syscall.SIGPIPE:
	default:
		fmt.Println("[Client] 收到信号准备退出...")
	}
	return tag
}
