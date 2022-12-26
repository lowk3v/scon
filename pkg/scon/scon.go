package scon

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/lowk3v/scon/config"
	"github.com/lowk3v/scon/internal/model"
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
			_, _ = fmt.Fprintf(os.Stderr, "%s %s is not a blockchain address", config.Symbol.Error, sc.Address)
			continue
		}

		go func() {
			defer wg.Done()

			if err := DetectChain(args.ChainNameOpt, sc); err != nil {
				if err.Error() == "not found" {
					_, _ = fmt.Fprintf(os.Stderr, "%s %s: %s\n",
						config.Symbol.Error,
						sc.Address,
						"Unknown")
				} else {
					_, _ = fmt.Fprintf(os.Stderr, "%s Error: %v\n", config.Symbol.Error, err)
				}
				return
			}
			if args.DetectChainMode {
				return
			}

			err := DumpSource(sc, args.OutputOpt)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%s Error: %v\n", config.Symbol.Error, err.Error())
			}
			if args.DumpSourceMode {
				return
			}
		}()

		wg.Wait()
	}
}
