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
			//var buffer bytes.Buffer
			//for i := 0; i < 1024; i++ {
			//	buffer.Write([]byte("test"))
			//}
			//
			//t.Log("finish")
			//t.Log(len(buffer.Bytes()))
			//t.Log(binary.Size(buffer.Bytes()))
			client.StartConn()
		})
	}
}
