package storage

import (
	bk "github.com/blockchain-db/core/common"
	"time"
	"github.com/blockchain-db/util"
	"errors"
)

type InMemoryBlockStore struct {
	blockMap map[string]bk.Block
	blocks []string
	blockChain *bk.BlockChain
}

func NewInMemoryBlockStore() *InMemoryBlockStore {
	//Create new block store with default configuration
	blockStore := &InMemoryBlockStore{}
	blockStore.blockMap = make(map[string]bk.Block)
	blockStore.blockChain = &bk.BlockChain{}
	blockStore.blockChain.Height = 0
	blockStore.blockChain.CurrentBlockHash = "0"

	//Add initial block to blockchain
	blockStore.AddBlock(getGenesisBlock())

	return blockStore
}

func getGenesisBlock() *bk.Block {
	gensisBlock := &bk.Block{}
	gensisBlock.Data = "Initial Block"

	return gensisBlock
}

func (blockStore *InMemoryBlockStore) AddBlock(block *bk.Block) error {
	blockMap := blockStore.blockMap
	blockChain := blockStore.blockChain

	//set fields in block
	block.Header = &bk.BlockHeader{}
	block.Header.Timestamp = time.Now().UnixNano()
	block.Header.Number = blockChain.Height
	hash := util.Hash(block.Data, block.Header.Timestamp)
	block.Header.DataHash = hash
	block.Header.PreviousHash = blockChain.CurrentBlockHash

	//Add block to chain
	blockMap[hash] = *block
	blockStore.blocks = append(blockStore.blocks, hash)

	//update blockchain information
	blockChain.Height += 1
	blockChain.PreviousBlockHash = blockChain.CurrentBlockHash
	blockChain.CurrentBlockHash = hash

	return nil
}


func (blockStore *InMemoryBlockStore) GetBlockByHash(blockHash []byte) (bk.Block, error) {
	var block = bk.Block{}
	if block, ok := blockStore.blockMap[string(blockHash)]; ok {
		return block, nil
	}
	return block, errors.New("Block with this hash doesn't exist")
}


func (blockStore *InMemoryBlockStore) GetBlockByNumber(blockNum uint64) (bk.Block, error) {
	var block = bk.Block{}
	if blockNum >= uint64(len(blockStore.blocks)) {
		return block, errors.New("Block number doesn't exist")
	}
	hash := blockStore.blocks[blockNum]
	block = blockStore.blockMap[hash]
	return block, nil
}

