package msg

import "github.com/ioeXNetwork/ioeX.Utility/p2p"

// Ensure Pong implement p2p.Message interface.
var _ p2p.Message = (*Pong)(nil)

type Pong struct {
	Ping
}

func NewPong(nonce uint64) *Pong {
	pong := new(Pong)
	pong.Nonce = nonce
	return pong
}

func (msg *Pong) CMD() string {
	return p2p.CmdPong
}
