package scon

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/fatih/color"

	"github.com/lowk3v/scon/config"
	"github.com/lowk3v/scon/internal/bsc"
	"github.com/lowk3v/scon/internal/model"
	"github.com/lowk3v/scon/internal/utils/file"
)

func DumpSource(sc *model.SmartContract, outputDir string) error {
	if !sc.IsValidAddress() || !sc.HasChain() {
		return errors.New("token is not valid or chain not found")
	}
	if err := file.Mkdir(outputDir); err != nil {
		return err
	}

	for idx, chain := range sc.Chains {
		switch chain.ChainId {
		case config.AppConfig.BscScan.ChainId:
			if err := bsc.DumpSource(sc.Address, &chain); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "%s Error: %v\n", config.Symbol.Error, err)
			}
		}
		sc.Chains[idx] = chain

		// write files
		for _, source := range chain.Contract.Sourcecode {
			subDir := filepath.Join(outputDir, sc.Address, chain.ChainName)
			fileSrcName := filepath.Join(subDir, source["filename"])
			if err := file.Mkdir(subDir); err != nil {
				return err
			}
			if err := file.WriteFile(fileSrcName, source["content"]); err != nil {
				return err
			}
			fmt.Printf("%s Dumped source code to: %s\n", config.Symbol.Success, color.BlueString(fileSrcName))
		}
	}

	return nil
}
