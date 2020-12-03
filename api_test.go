package wattility_go_sdk

import (
	"github.com/stretchr/testify/assert"
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
	//_ = client.client.SetProxy("http://localhost:9091")
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
			_, err := client.CheckSignApi()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestClient_LoadBaseSummary(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var data []LoadBaseSummaryBody
			data1 := LoadBaseSummaryBody{
				DatetimeStart: "1",
				DatetimeEnd:   "1",
				ValueBase:     0,
				ValueReal:     0,
				ValueWBase:    0,
				ValueWReal:    0,
				ValueHBase:    0,
				ValueHReal:    0,
			}
			data2 := LoadBaseSummaryBody{
				DatetimeStart: "2",
				DatetimeEnd:   "2",
				ValueBase:     0,
				ValueReal:     0,
				ValueWBase:    0,
				ValueWReal:    0,
				ValueHBase:    0,
				ValueHReal:    0,
			}
			data = append(data, data1)
			data = append(data, data2)

			err := client.LoadBaseSummary(data)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
