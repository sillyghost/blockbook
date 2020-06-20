package aryacoin

import (
	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

// magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0xe7eae188
)

// chain parameters
var (
	MainNetParams chaincfg.Params
	
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{23}
	MainNetParams.ScriptHashAddrID = []byte{50}
	
}

// LitecoinParser handle
type AryacoinParser struct {
	*btc.BitcoinParser
}

// NewLitecoinParser returns new LitecoinParser instance
func NewAryacoinParser(params *chaincfg.Params, c *btc.Configuration) *AryacoinParser {
	return &AryacoinParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Litecoin network,
// and the test Litecoin network
func GetChainParams(chain string) *chaincfg.Params {
	// register bitcoin parameters in addition to litecoin parameters
	// litecoin has dual standard of addresses and we want to be able to
	// parse both standards
	
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&MainNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &MainNetParams
	default:
		return &MainNetParams
	}
}
