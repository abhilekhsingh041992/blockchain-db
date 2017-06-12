package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	mux "github.com/gorilla/mux"
	pr "github.com/blockchain-db/core/peer"
	bk "github.com/blockchain-db/core/common"
	bks "github.com/blockchain-db/core/storage"
	"strconv"
	"io/ioutil"
)

var peerInfo *pr.PeerInfo
var blockStore *bks.InMemoryBlockStore


func handleRoutes(router *mux.Router) {

	peerInfo = &pr.PeerInfo{}
	peerInfo.Peers = []pr.Peer{}
	blockStore = bks.NewInMemoryBlockStore()

	router.HandleFunc("/", Index)
	router.HandleFunc("/peers", getPeers).Methods("GET")
	router.HandleFunc("/peers", AddPeer).Methods("POST")
	router.HandleFunc("/blocks", AddBlock).Methods("POST")
	router.HandleFunc("/blocks/{id}", GetBlockByNumber).Methods("GET")
	router.HandleFunc("/blocks/hash/{blockHash}", GetBlockByHash).Methods("GET")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	handleRoutes(router)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func getPeers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(peerInfo.Peers)
}

func AddPeer(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	address := vars["address"]
	port := vars["port"]
	peer, _ := peerInfo.AddPeer(address, port)
	json.NewEncoder(w).Encode(peer)
}

func AddBlock(w http.ResponseWriter, r *http.Request) {
	var block bk.Block
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &block)
	blockStore.AddBlock(&block)
	json.NewEncoder(w).Encode(block)
}

func GetBlockByNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockNum, _ := strconv.ParseUint(vars["id"], 10, 64)
	block, _ := blockStore.GetBlockByNumber(blockNum)
	json.NewEncoder(w).Encode(block)
}

func GetBlockByHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockHash := vars["blockHash"]
	block, _ := blockStore.GetBlockByHash([]byte(blockHash))
	json.NewEncoder(w).Encode(block)
}