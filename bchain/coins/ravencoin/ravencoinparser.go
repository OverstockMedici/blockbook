package ravencoin

import (
	"blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

const (
	MainnetMagic wire.BitcoinNet = 0xdbb6c0fb
	TestnetMagic wire.BitcoinNet = 0xf1c8d2fd
	RegtestMagic wire.BitcoinNet = 0xdab5bffa
)

var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{48}
	MainNetParams.ScriptHashAddrID = []byte{50}
	MainNetParams.Bech32HRPSegwit = "ltc"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{111}
	TestNetParams.ScriptHashAddrID = []byte{58}
	TestNetParams.Bech32HRPSegwit = "tltc"
}

// ravencoinParser handle
type ravencoinParser struct {
	*btc.BitcoinParser
}

// NewravencoinParser returns new ravencoinParser instance
func NewravencoinParser(params *chaincfg.Params, c *btc.Configuration) *ravencoinParser {
	return &ravencoinParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main ravencoin network,
// and the test ravencoin network
func GetChainParams(chain string) *chaincfg.Params {
	// register bitcoin parameters in addition to ravencoin parameters
	// ravencoin has dual standard of addresses and we want to be able to
	// parse both standards
	if !chaincfg.IsRegistered(&chaincfg.MainNetParams) {
		chaincfg.RegisterBitcoinParams()
	}
	if !chaincfg.IsRegistered(&MainNetParams) {
		err := chaincfg.Register(&MainNetParams)
		if err == nil {
			err = chaincfg.Register(&TestNetParams)
		}
		if err != nil {
			panic(err)
		}
	}
	switch chain {
	case "test":
		return &TestNetParams
	default:
		return &MainNetParams
	}
}
