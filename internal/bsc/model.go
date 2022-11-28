package bsc

import (
	"errors"
	"fmt"

	"github.com/tidwall/gjson"

	"scon/internal/model"
)

type ApiResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Raw     string
	Data    []model.Contract
}

func (c *ApiResponse) parseFromJson(json string) error {
	raw := gjson.Parse(json)
	c.Status = raw.Get("status").String()
	c.Message = raw.Get("message").String()

	if c.Status == "0" {
		// todo check contract is verified?
		return nil
	}

	for _, result := range raw.Get("result").Array() {

		cont := model.Contract{
			Abi:            result.Get("ABI").String(),
			ContractName:   result.Get("ContractName").String(),
			Proxy:          result.Get("Proxy").String(),
			Implementation: result.Get("Implementation").String(),
			SwarmSource:    result.Get("SwarmSource").String(),
		}

		if result.Get("SourceCode").String() == "" {
			return errors.New("contract source code not verified")
		}

		source, isOk := gjson.Parse(result.Get("SourceCode").String()).Value().(map[string]interface{})
		if !isOk {
			return errors.New(fmt.Sprintf("cannot parse source code: %v", raw.Get("SourceCode").String()))
		}

		for k, v := range source {
			v2 := v.(map[string]interface{})
			cont.Sourcecode = append(cont.Sourcecode, map[string]string{
				"filename": k,
				"content":  v2["content"].(string),
			})
		}
		c.Data = append(c.Data, cont)
	}

	return nil
}
