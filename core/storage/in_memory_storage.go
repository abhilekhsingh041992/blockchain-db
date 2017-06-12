package storage

import (
	bk "github.com/blockchain-db/core/common"
)

type InMemoryBlockStore struct {
	blocks map[string]bk.Block
	blockChain *bk.BlockChain
}

func NewInMemoryBlockStore() *InMemoryBlockStore {
	blockStore := &InMemoryBlockStore{}
	blockStore.blocks = make(map[string]bk.Block)
	blockStore.blockChain = &bk.BlockChain{}
	blockStore.blockChain.Height = 0
	return blockStore
}

func (blockStore *InMemoryBlockStore) AddBlock(block *bk.Block) error {
	blocks := blockStore.blocks
	blockChain := blockStore.blockChain

	block.Header.Number = blockChain.Height
	blockChain.Height += 1
	hash := block.Header.DataHash

	blocks[string(hash)] = *block
	blockChain.PreviousBlockHash = blockChain.CurrentBlockHash
	blockChain.CurrentBlockHash = hash
	return nil
}


func (blockStore *InMemoryBlockStore) GetBlockByHash(blockHash []byte) (bk.Block, error) {
	return blockStore.blocks[string(blockHash)], nil
}


func (blockStore *InMemoryBlockStore) GetBlockByNumber(blockNum uint64) (bk.Block, error) {
	block := bk.Block{}
	return block, nil
}

