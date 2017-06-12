package main

import (
	"fmt"
	"log"
	"net/http"
	mux "github.com/gorilla/mux"
	pr "github.com/blockchain-db/core/peer"
	bks "github.com/blockchain-db/core/storage"
	state "github.com/blockchain-db/core/state"
	route "github.com/blockchain-db/api/route"

)


func handleRoutes(router *mux.Router) {
	state.PeerInfo = &pr.PeerInfo{}
	state.PeerInfo.Peers = []pr.Peer{}
	state.BlockStore = bks.NewInMemoryBlockStore()

	router.HandleFunc("/", Index)
	router.HandleFunc("/peers", route.GetPeers).Methods("GET")
	router.HandleFunc("/peers", route.AddPeer).Methods("POST")
	router.HandleFunc("/blocks", route.AddBlock).Methods("POST")
	router.HandleFunc("/blocks/{id}", route.GetBlockByNumber).Methods("GET")
	router.HandleFunc("/blocks/hash/{blockHash}", route.GetBlockByHash).Methods("GET")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	handleRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}
