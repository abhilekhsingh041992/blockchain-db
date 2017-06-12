package routes

import (

)


type PeerInfo struct {
	Peers []Peer
}

type Peer struct {
	address string
	port string
}

func (peerInfo *PeerInfo) GetPeers() []Peer {
	return peerInfo.Peers
}

func (peerInfo *PeerInfo) AddPeer(address string, port string) (*Peer, error) {
	peer := &Peer{}
	peer.address = address
	peer.port = port
	peerInfo.Peers = append(peerInfo.Peers, *peer)
	return peer, nil
}