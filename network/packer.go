package network

import (
	"encoding/binary"
	"io"
	"net"
	"time"
)

type NormalPacker struct {
	ByteOrder binary.ByteOrder
}

func NewNormalPacker(order binary.ByteOrder) *NormalPacker {
	return &NormalPacker{ByteOrder: order}
}

// Pack |data 长度|id|data
func (p *NormalPacker) Pack(message *Message) ([]byte, error) {
	buffer := make([]byte, 8+8+len(message.Data))
	p.ByteOrder.PutUint64(buffer[0:8], uint64(len(buffer))) // total length
	p.ByteOrder.PutUint64(buffer[8:16], message.Id)
	copy(buffer[16:], message.Data)
	return buffer, nil
}

// UnPack returns Message
func (p *NormalPacker) UnPack(reader io.Reader) (*Message, error) {
	err := reader.(*net.TCPConn).SetReadDeadline(time.Now().Add(time.Second * 10))
	if err != nil {
		return nil, err
	}
	buffer := make([]byte, 8+8)
	if _, err := io.ReadFull(reader, buffer); err != nil {
		return nil, err
	}
	totalLen := p.ByteOrder.Uint64(buffer[:8])
	id := p.ByteOrder.Uint64(buffer[8:])
	dataSize := totalLen - 8 - 8
	data := make([]byte, dataSize)
	if _, err := io.ReadFull(reader, data); err != nil {
		return nil, err
	}
	msg := &Message{
		Id:   id,
		Data: data,
	}
	return msg, nil
}
