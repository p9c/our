package mod

import (
	"github.com/parallelcointeam/our/amp"
)

type Block struct {
	BlockHeight int    `json:"blockheight"`
	BlockHash   string `json:"bhash"`
	Data        []byte `json:"b"`
}

type Blk struct {
	Block string `json:"b"`
}
type Tx struct {
	TxHash string `json:"txhash""`
	Data   []byte `json:"data"`
}

type Addr struct {
	Addr string `json:"addr"`
}

type Explorer struct {
	Coin     Coin     `json:"coin"`
	AMPblock amp.AMPb `json:"amp"`
}

type BlVw struct {
	ID       string   `json:"id"`
	Coin     Coin     `json:"coin"`
	Block    Block    `json:"block"`
	AMPblock amp.AMPb `json:"amp"`
}

type TxVw struct {
	ID       string   `json:"id"`
	Coin     Coin     `json:"coin"`
	Tx       Tx       `json:"tx"`
	AMPblock amp.AMPb `json:"amp"`
}

type AdVw struct {
	ID       string   `json:"id"`
	Coin     Coin     `json:"coin"`
	Addr     Addr     `json:"addr"`
	AMPblock amp.AMPb `json:"amp"`
}
