package scon

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lowk3v/scon/internal/model"
	"github.com/lowk3v/scon/pkg/scon"
)

func Test_detectChain(t *testing.T) {
	type args struct {
		sc     *model.SmartContract
		expect int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"addressIsBSC",
			args{
				&model.SmartContract{
					Address: "0xb3c1cfd7c0b34090f9500f8fbc1ed805629b8707",
				},
				1,
			},
		},
		{
			"addressNotBSC",
			args{
				&model.SmartContract{
					Address: "0xb3c1cfd7c0b34090f9500f8fbc1ed805629b8000",
				},
				0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := scon.DetectChain("", tt.args.sc)
			assert.NoError(t, err, fmt.Sprintf("Error: %s", err))
			assert.Equal(t, tt.args.expect, len(tt.args.sc.Chains), "")
		})
	}
}
