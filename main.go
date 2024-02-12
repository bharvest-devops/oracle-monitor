package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"bharvest.io/oraclemon/app"
	"bharvest.io/oraclemon/log"
	"bharvest.io/oraclemon/server"
	"bharvest.io/oraclemon/tg"
	"bharvest.io/oraclemon/wallet"
	"github.com/BurntSushi/toml"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	ctx := context.Background()

	cfgPath := flag.String("config", "", "Config file")
    flag.Parse()
	if *cfgPath == "" {
		panic("Error: Please input config file path with -config flag.")
	}

	f, err := os.ReadFile(*cfgPath)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	cfg := app.Config{}
	err = toml.Unmarshal(f, &cfg)
	if err != nil {
		log.Error(err)
		panic(err)
	}

	if cfg.UmeeOracle.Enable {
		err = PrePrepareForUmee(ctx, &cfg)
		if err != nil {
			log.Error(err)
			panic(err)
		}
	}

	tgTitle := fmt.Sprintf("🤖 Oracle Monitor for %s 🤖", cfg.General.Network)
	tg.SetTg(cfg.Tg.Enable, tgTitle, cfg.Tg.Token, cfg.Tg.ChatID)

	go server.Run(cfg.General.ListenPort)
	for {
		app.Run(ctx, &cfg)
		time.Sleep(time.Duration(cfg.General.Period) * time.Minute)
	}
}

func PrePrepareForUmee(ctx context.Context, cfg *app.Config) error {
	var err error
	cfg.UmeeOracle.Wallet, err = wallet.NewWallet(ctx, cfg.UmeeOracle.ValidatorAcc)
	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
