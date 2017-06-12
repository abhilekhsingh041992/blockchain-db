package storage

import (
	bk "github.com/blockchain-db/core/common"
)

type BlockStore interface {

	AddBlock(block *bk.Block) error
	GetBlockByHash(blockHash []byte) (bk.Block, error)
	GetBlockByNumber(blockNum uint64) (bk.Block, error)

}
