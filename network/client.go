package network

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

type Client struct {
	address string
	network string
	packer  IPacker
	chMsg   chan *Message
}

func NewClient(address, network string) *Client {
	return &Client{
		address: address,
		network: network,
		packer: &NormalPacker{
			ByteOrder: binary.BigEndian,
		},
		chMsg: make(chan *Message, 1),
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

func (c Client) Write(conn net.Conn) (*Message, error) {
	tick := time.NewTicker(time.Second)
	for {
		select {
		case <-tick.C:
			c.chMsg <- &Message{
				Id:   11111,
				Data: []byte("Hello World!"),
			}
		case msg := <-c.chMsg:
			c.send(conn, msg)
		}
	}
}

func (c Client) Read(conn net.Conn) {
	for {
		message, err := c.packer.UnPack(conn)
		if _, ok := err.(net.Error); err != nil && ok { // 处理网络原因
			fmt.Println(err)
			continue
		}
		fmt.Println("resp message:", string(message.Data))
	}
}

func (c Client) send(conn net.Conn, message *Message) {
	pack, err := c.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	conn.Write(pack)
}
