package app

import (
	"context"
	"sync"
	"time"

	"bharvest.io/oraclemon/log"
)

func Run(ctx context.Context, cfg *Config) {
	ctx, cancel := context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)

	// Check Umee Oracle
	go func() {
		defer wg.Done()
		if !cfg.UmeeOracle.Enable {
			return
		}

		err := cfg.CheckUmeeOracle(ctx)
		if err != nil {
			log.Error(err)
			return
		}
	}()

	wg.Wait()

	return
}
