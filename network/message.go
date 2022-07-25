package network

type Message struct {
	Id   uint64
	Data []byte // 方便转码
}
