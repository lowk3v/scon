package bsc

import (
	"errors"
	"fmt"
	"strings"

	"scon/config"
	"scon/internal/model"
	"scon/internal/utils"
)

func IsBscChain(address string) (bool, error) {
	api := fmt.Sprintf("%s%s",
		config.AppConfig.BscScan.Host,
		config.AppConfig.BscScan.SearchHandler)
	api = strings.ReplaceAll(api, "$ADDRESS", address)

	resp, err := utils.HttpGet(api)
	if err != nil {
		return false, err
	}

	if strings.Contains(
		strings.ToLower(resp),
		strings.ToLower(address)) {
		return true, nil
	}
	return false, errors.New("the address not found")
}

func DumpSource(address string, chain *model.Chain) error {
	api := fmt.Sprintf("%s%s",
		config.AppConfig.BscScan.Api,
		config.AppConfig.BscScan.GetSourcecode)
	api = strings.ReplaceAll(api, "$ADDRESS", address)
	api = strings.ReplaceAll(api, "$APIKEY", config.Secret.BscScanKey)

	respRaw, err := utils.HttpGet(api)
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
