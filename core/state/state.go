package state

import (
	pr "github.com/blockchain-db/core/peer"
	bks "github.com/blockchain-db/core/storage"
)

var (
	PeerInfo *pr.PeerInfo
	BlockStore *bks.InMemoryBlockStore
)

