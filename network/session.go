package network

import (
	"encoding/binary"
	"github.com/lqgl/tinywork/logger"
	"net"
	"time"
)

type Session struct {
	UId            uint64
	conn           net.Conn
	IsClose        bool
	packer         *NormalPacker
	ChWrite        chan *Message
	IsPlayerOnline bool
	MessageHandler func(packet *SessionPacket)
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		conn:    conn,
		packer:  NewNormalPacker(binary.BigEndian),
		ChWrite: make(chan *Message, 10),
	}
}

func (s *Session) Run() {
	go s.Read()
	go s.Write()
}

// Read 接收来自客户端的消息
func (s *Session) Read() {
	for {
		err := s.conn.SetReadDeadline(time.Now().Add(time.Second)) // 设置读字节流超时时间
		if err != nil {
			logger.Logger.DebugF("%v\n", err)
			continue
		}
		message, err := s.packer.UnPack(s.conn)
		if _, ok := err.(net.Error); ok { // 处理网络原因
			continue
		}
		s.MessageHandler(&SessionPacket{ // 将收到的消息发送给处理路由
			Msg:  message,
			Sess: s,
		})
	}
}

// Write 返回数据
func (s *Session) Write() {
	for {
		select {
		case resp := <-s.ChWrite:
			s.send(resp)
		}
	}
}

func (s *Session) send(message *Message) {
	err := s.conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		logger.Logger.DebugF("%v\n", err)
		return
	}
	bytes, err := s.packer.Pack(message)
	if err != nil {
		logger.Logger.DebugF("%v\n", err)
		return
	}
	s.conn.Write(bytes)
}

func (s *Session) SendMsg(msg *Message) {
	s.ChWrite <- msg
}
