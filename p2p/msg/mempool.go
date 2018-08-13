package msg

import (
	"io"

	"github.com/ioeX/ioeX.Utility/p2p"
)

type MemPool struct{}

func (msg *MemPool) CMD() string {
	return p2p.CmdMemPool
}

func (msg *MemPool) Serialize(io.Writer) error {
	return nil
}

func (msg *MemPool) Deserialize(io.Reader) error {
	return nil
}
