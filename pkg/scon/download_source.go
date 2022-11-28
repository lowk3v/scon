package scon

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"

	"scon/config"
	"scon/internal/bsc"
	"scon/internal/model"
	"scon/internal/utils"
)

func DumpSource(sc *model.SmartContract, outputDir string) error {
	if !sc.IsValidAddress() || !sc.HasChain() {
		return errors.New("token is not valid or chain not found")
	}
	if err := utils.Mkdir(outputDir); err != nil {
		return err
	}

	for idx, chain := range sc.Chains {
		switch chain.ChainId {
		case config.AppConfig.BscScan.ChainId:
			if err := bsc.DumpSource(sc.Address, &chain); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
		}
		sc.Chains[idx] = chain

		// write files
		for _, source := range chain.Contract.Sourcecode {
			subDir := filepath.Join(outputDir, sc.Address, chain.ChainName)
			if err := utils.Mkdir(subDir); err != nil {
				return err
			}
			if err := utils.WriteFile(filepath.Join(subDir, source["filename"]), source["content"]); err != nil {
				return err
			}
		}
		fmt.Printf("[+] Dumped source code to a directory: %s\n", color.BlueString(outputDir))

	}

	return nil
}
