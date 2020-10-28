package wattility_go_sdk

import (
	"testing"
)

func TestClient_do(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient("aaaaaaaaaaaaaaaa", "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa")
			c.Client.SetProxy("http://localhost:9091")
			c.do("/test?foo=bar", "POST", []byte(`{"username":"testuser", "password":"testpass"}`))
		})
	}
}
