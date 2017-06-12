package route

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	state "github.com/blockchain-db/core/state"
)

func GetPeers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(state.PeerInfo.GetPeers())
}

func AddPeer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	port := vars["port"]
	peer, _ := state.PeerInfo.AddPeer(address, port)
	json.NewEncoder(w).Encode(peer)
}
