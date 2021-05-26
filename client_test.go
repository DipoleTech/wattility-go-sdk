package wattility_go_sdk

import (
	"testing"
)

var (
	tcpAddr    = "127.0.0.1:8999"
	testKey    = "zp7mOixKK8Sd0O7M"
	testSecret = "qiHRQerjVkcGIyoDEi0phCsoxhdODUAF"
)

func TestClient_StartConn(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(testKey, testSecret)
			c.SetHost(tcpAddr)
			c.SetDebug()
			c.StartConn()
		})
	}
}
