package model

import (
	"regexp"
)

type SmartContract struct {
	Address string
	Chains  []Chain
}

type Chain struct {
	ChainId   int
	ChainName string
	TokenName string
	Contract  Contract
}

type Contract struct {
	Sourcecode     []map[string]string
	Abi            string
	ContractName   string
	Proxy          string
	Implementation string
	SwarmSource    string
}

func (sc *SmartContract) IsValidAddress() bool {
	var pattern = regexp.MustCompile("^0x[0-9a-zA-Z]{40}$")
	if sc.Address == "" {
		return false
	}
	if !pattern.MatchString(sc.Address) {
		return false
	}
	return true
}

func (sc *SmartContract) HasChain() bool {
	return len(sc.Chains) > 0
}
