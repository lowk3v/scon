package scon

import (
	"fmt"
	"strings"

	"github.com/fatih/color"

	"github.com/lowk3v/scon/config"
	"github.com/lowk3v/scon/internal/bsc"
	"github.com/lowk3v/scon/internal/model"
)

func DetectChain(inputChainName string, sc *model.SmartContract) error {
	if inputChainName != "" {
		passedFromArg(sc, strings.Split(inputChainName, ","))
	}

	if !sc.IsValidAddress() {
		return fmt.Errorf("invalid address")
	}

	if isBscChain, err := bsc.IsBscChain(sc.Address); err != nil {
		return err
	} else if isBscChain {
		newChain := model.Chain{
			ChainId:   config.AppConfig.BscScan.ChainId,
			ChainName: config.AppConfig.BscScan.ChainName,
		}
		sc.Chains = append(sc.Chains, newChain)
		fmt.Printf("%s %s: %s[%d]\n",
			config.Symbol.Success,
			sc.Address,
			color.BlueString(newChain.ChainName),
			newChain.ChainId)
	}
	return nil
}

func passedFromArg(sc *model.SmartContract, chainsName []string) {
	for _, chainName := range chainsName {
		switch strings.ToLower(chainName) {
		case strings.ToLower(config.AppConfig.BscScan.ChainName):
			sc.Chains = append(sc.Chains, model.Chain{
				ChainId:   config.AppConfig.BscScan.ChainId,
				ChainName: config.AppConfig.BscScan.ChainName,
			})
		}

	}
}
