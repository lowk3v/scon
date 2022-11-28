package scon

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"scon/config"
	"scon/internal/model"
	"scon/pkg/scon"
)

func init() {
	config.InitHttpClient("")
}

func TestDownloadSource(t *testing.T) {
	type args struct {
		sc        *model.SmartContract
		outputDir string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"NotVerified",
			args{
				sc: &model.SmartContract{
					Address: "0xb3c1cfd7c0b34090f9500f8fbc1ed805629b8707",
					Chains: []model.Chain{
						{
							ChainId:   config.AppConfig.BscScan.ChainId,
							ChainName: config.AppConfig.BscScan.ChainName,
						},
					},
				},
				outputDir: "./testdata",
			},
			true,
		},
		{
			"Verified",
			args{
				sc: &model.SmartContract{
					Address: "0x61AA9b7e74104c74b0231aF3641A6a804E81761F",
					Chains: []model.Chain{
						{
							ChainId:   config.AppConfig.BscScan.ChainId,
							ChainName: config.AppConfig.BscScan.ChainName,
						},
					},
				},
				outputDir: "./testdata",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := scon.DumpSource(tt.args.sc, tt.args.outputDir); (err != nil) != tt.wantErr {
				if tt.wantErr {
					assert.NoError(t, err)
				} else {
					assert.Error(t, err, fmt.Sprintf("Error: %v", err))
				}
			}
		})
	}
}
