package wattility_go_sdk

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	client    *Client
	appId     = "aaaaaaaaaaaaaaaa"
	appSecret = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"
)

func init() {
	var err error
	client, err = NewClient(appId, appSecret)

	if err != nil {
		panic(err)
	}
	_ = client.client.SetProxy("http://localhost:9091")
}

func TestClient_CheckAccessApi(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		{"test", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.CheckAccessAPI()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}
