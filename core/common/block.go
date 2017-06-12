package common

import (
	"encoding/json"
	"github.com/blockchain-db/util"
)

type Block struct {
	Header *BlockHeader  `json:"header"`
	Data string  `json:"data"`
	MetaData *BlockMetaData  `json:"metaData"`
}


type BlockHeader struct {
	Number       uint64  `json:"number"`
	Timestamp    int64  `json:"timestamp"`
	PreviousHash string  `json:"previousHash"`
	DataHash     string  `json:"dataHash"`
}


type BlockMetaData struct {
	Metadata []string  `json:"metaData"`
}


func NewBlock() *Block {
	block := &Block{}
	block.Header = &BlockHeader{}
	block.Header.Number = 1

	return block
}

func (block *Block) Hash() (string, error)  {
	dataBytes, err := json.Marshal(block)
	if err != nil {
		return "", nil
	}
	return util.Hash(string(dataBytes)), nil
}