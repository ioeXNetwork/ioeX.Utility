package msg

import (
	"github.com/ioeXNetwork/ioeX.Utility/p2p"
)

type NotFound struct {
	Inventory
}

func NewNotFound() *NotFound {
	msg := &NotFound{
		Inventory: *NewInventory(),
	}
	return msg
}

func (msg *NotFound) CMD() string {
	return p2p.CmdNotFound
}
