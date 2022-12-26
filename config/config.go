package config

import (
	"crypto/tls"
	_ "embed"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
)

type secretConfig struct {
	EnvMode    string `json:"SCON_ENV"`
	BscScanKey string `json:"SCON_BSC_SCAN_KEY"`
}

type appConfig struct {
	BscScan ExplorerConfig `json:"bscscan"`
}

type ExplorerConfig struct {
	ChainId          int    `json:"chain-id"`
	ChainName        string `json:"chain-name"`
	Api              string `json:"api"`
	Host             string `json:"host"`
	TokenInfo        string `json:"token-info"`
	SearchHandler    string `json:"search-handler"`
	TokenTotalSupply string `json:"token-total-supply"`
	GetSourcecode    string `json:"get-sourcecode"`
}

type SymbolConfig struct {
	Success string
	Error   string
	Info    string
}

var Secret secretConfig
var AppConfig appConfig
var HttpClient *http.Client
var Symbol SymbolConfig

//go:embed config.yaml
var appConfigYaml string

//go:embed .env
var envFile string

func init() {
	// load environment from file. Be overridden by system environment
	envContent, err := godotenv.Unmarshal(envFile)
	err = convertToStruct(envContent, &Secret)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s Error: %v\n", color.RedString("¿"), err)
	}
	loadSystemEnv(&Secret)

	// load config based on environment (mainnet, testnet)
	var appConfigTemp map[string]interface{}
	err = yaml.Unmarshal([]byte(appConfigYaml), &appConfigTemp)
	err = convertToStruct(appConfigTemp[Secret.EnvMode], &AppConfig)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s Error: %v\n", color.RedString("¿"), err)
	}

	Symbol = SymbolConfig{
		Success: color.GreenString("≠"),
		Error:   color.RedString("¿"),
		Info:    color.BlueString("ℹ"),
	}
}

func loadSystemEnv(s *secretConfig) {
	if len(os.Getenv("SCON_ENV")) != 0 {
		s.EnvMode = os.Getenv("SCON_ENV")
	}
	if len(os.Getenv("SCON_BSC_SCAN_KEY")) != 0 {
		s.BscScanKey = os.Getenv("SCON_BSC_SCAN_KEY")
	}
}

func convertToStruct(from interface{}, to interface{}) error {
	jsonS, err := json.Marshal(from)
	err = json.Unmarshal(jsonS, to)
	if err != nil {
		return err
	}
	return nil
}

func InitHttpClient(proxy string) {
	transport := &http.Transport{
		MaxIdleConns:    30,
		IdleConnTimeout: time.Second,
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		DialContext: (&net.Dialer{
			Timeout:   time.Second * 10,
			KeepAlive: time.Second,
		}).DialContext,
	}

	if proxy != "" {
		if p, err := url.Parse(proxy); err == nil {
			transport.Proxy = http.ProxyURL(p)
		}
	}

	redirect := func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	HttpClient = &http.Client{
		Transport:     transport,
		CheckRedirect: redirect,
		Timeout:       time.Second * 10,
	}
}

func (app *appConfig) SupportedChains() {
	fmt.Printf("%s Supported Chains: %s[%d]\n",
		color.GreenString("≠"),
		color.BlueString(app.BscScan.ChainName),
		app.BscScan.ChainId)
}

func (s *secretConfig) PrintInfo() {
	fmt.Printf("%s Switching to [%s]. To run testnet set SCON_ENV=testnet\n",
		Symbol.Info,
		color.BlueString(Secret.EnvMode))
	if s.BscScanKey != "" {
		fmt.Printf("%s Loaded BscScan API Key: %s\n",
			Symbol.Info,
			color.BlueString("***"))
	}
	fmt.Println("")
}
