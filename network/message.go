package network

import "github.com/bishalbera/modular-blockchain/core"

type GetBlocksMessage struct {
	From uint32
	// If To is 0 then the max blocks will be returned.
	To uint32
}

type BlocksMessage struct {
	Blocks []*core.Block
}

type GetStatusMessage struct{}

type StatusMessage struct {
	// id of the server
	ID            string
	Version       uint32
	CurrentHeight uint32
}
