package network

import (
	"fmt"
	"net"
)

type Server struct {
	tcpListener     net.Listener
	address         string
	network         string
	OnSessionPacket func(packet *SessionPacket)
}

func NewServer(address, network string) *Server {
	s := &Server{
		address: address,
		network: network,
	}
	return s
}

func (s *Server) Run() {
	resolveTCPAddr, err := net.ResolveTCPAddr(s.network, s.address)
	if err != nil {
		fmt.Println(err)
		return
	}
	tcpListener, err := net.ListenTCP(s.network, resolveTCPAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	s.tcpListener = tcpListener
	for {
		conn, err := s.tcpListener.Accept()
		if err != nil {
			continue
		}
		go func() {
			newSession := NewSession(conn)
			newSession.MessageHandler = s.OnSessionPacket // 初始化 session.MessageHandler
			SessionMgrInstance.AddSession(newSession)
			newSession.Run()
			SessionMgrInstance.DelSession(newSession.UId)
		}()
	}
}
