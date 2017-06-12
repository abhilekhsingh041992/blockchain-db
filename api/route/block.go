package route

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	 bk "github.com/blockchain-db/core/common"
	 state "github.com/blockchain-db/core/state"
)

func AddBlock(w http.ResponseWriter, r *http.Request) {
	var block bk.Block
	body, _ := ioutil.ReadAll(r.Body)
	json.Unmarshal(body, &block)
	state.BlockStore.AddBlock(&block)
	json.NewEncoder(w).Encode(block)
}

func GetBlockByNumber(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockNum, _ := strconv.ParseUint(vars["id"], 10, 64)
	block, _ := state.BlockStore.GetBlockByNumber(blockNum)
	json.NewEncoder(w).Encode(block)
}

func GetBlockByHash(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockHash := vars["blockHash"]
	block, _ := state.BlockStore.GetBlockByHash([]byte(blockHash))
	json.NewEncoder(w).Encode(block)
}
