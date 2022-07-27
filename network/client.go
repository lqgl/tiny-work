package network

import (
	"encoding/binary"
	"fmt"
	"net"
)

type Client struct {
	address   string
	network   string
	packer    IPacker
	ChMsg     chan *Message
	OnMessage func(message *ClientPacket)
}

func NewClient(address, network string) *Client {
	return &Client{
		address: address,
		network: network,
		packer: &NormalPacker{
			ByteOrder: binary.BigEndian,
		},
		ChMsg: make(chan *Message, 1),
	}
}

func (c Client) Run() {
	conn, err := net.Dial(c.network, c.address)
	if err != nil {
		fmt.Println(err)
		return
	}

	go c.Read(conn)
	go c.Write(conn)
}

// Write 将收到的消息写入连接
func (c *Client) Write(conn net.Conn) {
	for {
		select {
		case msg := <-c.ChMsg: // cli 接收到来自 Client 的消息
			c.Send(conn, msg)
		}
	}
}

// Read 读取服务器返回的消息
func (c *Client) Read(conn net.Conn) {
	for {
		message, err := c.packer.UnPack(conn)
		if _, ok := err.(net.Error); err != nil && ok { // 处理网络原因
			fmt.Println(err)
			continue
		}
		c.OnMessage(&ClientPacket{
			Msg:  message,
			Conn: conn,
		})
	}
}

// Send 发送消息
func (c *Client) Send(conn net.Conn, message *Message) {
	pack, err := c.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write(pack)
}
