package proto

import "encoding/json"

//import common "github.com/tendermint/tendermint/libs/common"

// 用于解析证实信息
// Use to parse the attest information from the broadcast.
type AttestInformation struct {
	Action  string          `json:"action"`
	Ranking [][]interface{} `json:"ranking"`
}

type ranking struct {
	Address string `json:"attestee"`
}

type Ranking struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data []ranking `json:"result"`
}
//http://192.168.1.221:46657/tri_net_info
type TrustQueryReq struct {
	TxSignReq string `json:"tx"`

}
type TrustDataRespInfo struct {
	Address []string `json:"attestee"`

}
type RPCResponse struct {
	JSONRPC string          `json:"jsonrpc"`
	ID      string       `json:"id"`
	CODE    int             `json:"code"`
	//Result	 []byte `json:"result,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
	//Result  interface{} `json:"result,omitempty"`
	//Result  json.RawMessage `json:"result,omitempty"`
	//Error   *RPCError       `json:"error,omitempty"`
type ResultBroadcastTxCommit struct {
	//CheckTx   abci.ResponseCheckTx   `json:"check_tx"`
	CheckTx   ResponseCheckTx   `json:"check_tx"`
	DeliverTx ResponseDeliverTx `json:"deliver_tx"`
	Hash      []byte           `json:"hash"`
	Height    int64                  `json:"height"`