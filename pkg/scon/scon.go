package scon

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"scon/config"
	"scon/internal/model"
)

type Arguments struct {
	DelayOpt     int
	ChainNameOpt string
	OutputOpt    string
	EnvOpt       string
	ProxyOpt     string

	DetectChainMode    bool
	DumpSourceMode     bool
	SupportedChainMode bool
}

func Run(args *Arguments) {
	config.InitHttpClient(args.ProxyOpt)

	if args.SupportedChainMode {
		config.AppConfig.SupportedChains()
		return
	}

	var wg sync.WaitGroup
	stdin := bufio.NewScanner(os.Stdin)

	for stdin.Scan() {
		wg.Add(1)
		time.Sleep(time.Duration(args.DelayOpt * 1000000))

		sc := &model.SmartContract{
			Address: stdin.Text(),
		}

		if !sc.IsValidAddress() {
			_, _ = fmt.Fprintf(os.Stderr, "%s is not a blockchain address", sc.Address)
			continue
		}

		go func() {
			defer wg.Done()

			if err := DetectChain(args.ChainNameOpt, sc); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				return
			}
			if args.DetectChainMode {
				return
			}

			err := DumpSource(sc, args.OutputOpt)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, err.Error())
			}
			if args.DumpSourceMode {
				return
			}
		}()

		wg.Wait()
	}
}
