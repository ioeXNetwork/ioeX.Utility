package v0

import (
	"io"

	"github.com/ioeX/ioeX.Utility/common"
	"github.com/ioeX/ioeX.Utility/p2p"
)

type GetData struct {
	Hash common.Uint256
}

func NewGetData(hash common.Uint256) *GetData {
	return &GetData{Hash: hash}
}

func (msg *GetData) CMD() string {
	return p2p.CmdGetData
}

func (msg *GetData) Serialize(w io.Writer) error {
	return common.WriteElement(w, msg.Hash)
}

func (msg *GetData) Deserialize(r io.Reader) error {
	return common.ReadElement(r, &msg.Hash)
}
