package wattility_go_sdk

import (
	"testing"
)

func TestClient_StartConn(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client.StartConn()
		})
	}
}
