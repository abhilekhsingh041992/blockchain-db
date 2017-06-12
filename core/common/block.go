package common


type Block struct {
	Header *BlockHeader  `json:"header"`
	Data *string  `json:"data"`
	MetaData *BlockMetaData  `json:"metaData"`
}


type BlockHeader struct {
	Number       uint64  `json:"number"`
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
