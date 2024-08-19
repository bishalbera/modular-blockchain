package network

import "io"

type NetAddr string

type RPC struct {
	From    NetAddr
	Payload io.Reader
}

type Transport interface {
	Consume() <-chan RPC
	Connect(Transport) error
	SendMessage(NetAddr, []byte) error
	Addr() NetAddr
}
