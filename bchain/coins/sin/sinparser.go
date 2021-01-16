package sin

import (
	"blockbook/bchain/coins/btc"

	"github.com/martinboehm/btcd/wire"
	"github.com/martinboehm/btcutil/chaincfg"
)

// magic numbers
const (
	MainnetMagic wire.BitcoinNet = 0xb8d4ddf8 // sinovate mainnet
	TestnetMagic wire.BitcoinNet = 0xd8f4fdb8 // sinovate testnet
	RegtestMagic wire.BitcoinNet = 0xdab5bffa
)

// chain parameters
var (
	MainNetParams chaincfg.Params
	TestNetParams chaincfg.Params
)

func init() {
	MainNetParams = chaincfg.MainNetParams
	MainNetParams.Net = MainnetMagic
	MainNetParams.PubKeyHashAddrID = []byte{63}
	MainNetParams.ScriptHashAddrID = []byte{5}
	MainNetParams.Bech32HRPSegwit = "sin"

	TestNetParams = chaincfg.TestNet3Params
	TestNetParams.Net = TestnetMagic
	TestNetParams.PubKeyHashAddrID = []byte{63}
	TestNetParams.ScriptHashAddrID = []byte{5}
	TestNetParams.Bech32HRPSegwit = "tsin"
}

// SinovateParser handle
type SinovateParser struct {
	*btc.BitcoinParser
}

// NewSinovateParser returns new SinovateParser instance
func NewSinovateParser(params *chaincfg.Params, c *btc.Configuration) *SinovateParser {
	return &SinovateParser{BitcoinParser: btc.NewBitcoinParser(params, c)}
}

// GetChainParams contains network parameters for the main Sinovate network,
// and the test Sinovate network
func GetChainParams(chain string) *chaincfg.Params {
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
