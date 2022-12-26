package bsc

import (
	"errors"
	"fmt"
	"strings"

	"github.com/lowk3v/scon/config"
	"github.com/lowk3v/scon/internal/model"
	"github.com/lowk3v/scon/internal/utils/http"
)

func IsBscChain(address string) (bool, error) {
	api := fmt.Sprintf("%s%s",
		config.AppConfig.BscScan.Host,
		config.AppConfig.BscScan.SearchHandler)
	api = strings.ReplaceAll(api, "$ADDRESS", address)

	resp, err := http.HttpGet(api)
	if err != nil {
		return false, err
	}

	if strings.Contains(
		strings.ToLower(resp),
		strings.ToLower(address)) {
		return true, nil
	}
	return false, errors.New("not found")
}

func DumpSource(address string, chain *model.Chain) error {
	api := fmt.Sprintf("%s%s",
		config.AppConfig.BscScan.Api,
		config.AppConfig.BscScan.GetSourcecode)
	api = strings.ReplaceAll(api, "$ADDRESS", address)
	api = strings.ReplaceAll(api, "$APIKEY", config.Secret.BscScanKey)

	respRaw, err := http.HttpGet(api)
	if err != nil {
		return err
	}

	var resp ApiResponse
	if err = resp.parseFromJson(respRaw); err != nil {
		return err
	}

	if resp.Status != "1" && len(resp.Data) == 0 {
		return errors.New(resp.Message)
	}
	chain.Contract = resp.Data[0]
	return nil
}
