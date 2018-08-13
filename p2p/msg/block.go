package msg

import (
	"io"

	"github.com/ioeX/ioeX.Utility/common"
	"github.com/ioeX/ioeX.Utility/p2p"
)

type Block struct {
	Block common.Serializable
}

func NewBlock(block common.Serializable) *Block {
	return &Block{Block: block}
}

func (msg *Block) CMD() string {
	return p2p.CmdBlock
}

func (msg *Block) Serialize(writer io.Writer) error {
	return msg.Block.Serialize(writer)
}

func (msg *Block) Deserialize(reader io.Reader) error {
	return msg.Block.Deserialize(reader)
}
