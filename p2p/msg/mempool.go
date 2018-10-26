package msg

import (
	"github.com/ioeXNetwork/ioeX.Utility/p2p"
)

// Ensure MemPool implement p2p.Message interface.
var _ p2p.Message = (*MemPool)(nil)

type MemPool struct{ empty }

func (msg *MemPool) CMD() string {
	return p2p.CmdMemPool
}
