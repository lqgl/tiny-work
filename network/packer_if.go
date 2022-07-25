package network

import "io"

type IPacker interface {
	Pack(message *Message) ([]byte, error)
	UnPack(reader io.Reader) (*Message, error)
}
