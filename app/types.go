package app

import "bharvest.io/oraclemon/wallet"

type Config struct {
	General struct {
		Network    string `toml:"network"`
		EthAPI     string `toml:"eth_api"`
		ListenPort int    `toml:"listen_port"`
		Period     int    `toml:"period"`
	} `toml:"general"`
	UmeeOracle struct {
		Enable       bool    `toml:"enable"`
		GRPC         string  `toml:"grpc"`
		ValidatorAcc string  `toml:"validator_acc"`
		MinUptime    float64 `toml:"min_uptime"`
		Wallet       *wallet.Wallet
	} `toml:"umee-oracle"`
	Tg struct {
		Enable bool   `toml:"enable"`
		Token  string `toml:"token"`
		ChatID string `toml:"chat_id"`
	} `toml:"tg"`
}
