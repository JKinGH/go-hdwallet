package hdwallet

func init() {
	coins[BTC_TESTNET] = newBTCTestnet
}

type btc_testnet struct {
	*btc
}

func newBTCTestnet(key *Key) Wallet {
	BTCTestnet := newBTC(key).(*btc)
	BTCTestnet.name = "Bitcoin_Testnet"
	BTCTestnet.symbol = "BTC_TEST"
	BTCTestnet.key.opt.Params = &BTCTestnetParams

	return &btc_testnet{btc: BTCTestnet}
}