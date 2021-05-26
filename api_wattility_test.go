package wattility_go_sdk

import (
	"testing"
)

var (
	client    *Client
	appId     = "OxlPWsUPS1AFboMk"
	appSecret = "z48BT1uqa10HUQeNK87bfEjita1bnQ7i"
)

func init() {
	var err error
	client, err = NewClient(appId, appSecret)
	if err != nil {
		panic(err)
	}
	client.SetDebug()
}

func TestClient_CheckSignApiApi(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
		})
	}
}
