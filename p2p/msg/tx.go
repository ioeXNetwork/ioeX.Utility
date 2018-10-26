package msg

import (
	"github.com/ioeXNetwork/ioeX.Utility/common"
	"github.com/ioeXNetwork/ioeX.Utility/p2p"
)

// Ensure Tx implement p2p.Message interface.
var _ p2p.Message = (*Tx)(nil)

type Tx struct {
	common.Serializable
}

func NewTx(tx common.Serializable) *Tx {
	return &Tx{Serializable: tx}
}

func (msg *Tx) CMD() string {
	return p2p.CmdTx
}

func (msg *Tx) MaxLength() uint32 {
	return MaxBlockSize
}
