package scon

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lowk3v/scon/config"
)

func TestInit(t *testing.T) {
	tests := []struct {
		key string
	}{
		{},
	}

	for _, tcase := range tests {
		t.Run(tcase.key, func(t *testing.T) {
			assert.LessOrEqual(t, 1, len(config.Secret.BscScanKey), fmt.Sprintf("BscScan Key: %s\n", config.Secret.BscScanKey))
			assert.LessOrEqual(t, 1, len(config.AppConfig.BscScan.Api), fmt.Sprintf("BscScan Key: %s\n", config.AppConfig.BscScan.Api))
		})
	}
}
