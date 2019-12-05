package hdwallet

import "fmt"

func init() {
	coins[WICC_TESTNET] = newWICCTestnet
}

type wicc_testnet struct {
	name   string
	symbol string
	key    *Key
}

func newWICCTestnet(key *Key) Wallet {

	key.opt.Params = &WICCTestnetParams
	return &wicc_testnet{
		name:   "sdsd",
		symbol: "adaz",
		key:    key,
	}
}

func (c *wicc_testnet) GetType() uint32 {
	fmt.Println("c.key.opt.CoinType1=",c.key.opt.CoinType)
//	c.key.opt.CoinType = c.key.opt.CoinType - 1
	fmt.Println("c.key.opt.CoinType2=",c.key.opt.CoinType)
	return c.key.opt.CoinType
}

func (c *wicc_testnet) GetName() string {
	return c.name
}

func (c *wicc_testnet) GetSymbol() string {
	return c.symbol
}

func (c *wicc_testnet) GetKey() *Key {
	fmt.Println("c.key.opt.CoinType3=",c.key.opt.CoinType)
//	c.key.opt.CoinType = c.key.opt.CoinType - 1
	fmt.Println("c.key.opt.CoinType4=",c.key.opt.CoinType)
	return c.key
}

func (c *wicc_testnet) GetAddress() (string, error) {
	c.key.opt.CoinType = c.key.opt.CoinType - 1
	return c.key.AddressBTC()
}
