package wattility_go_sdk

import (
	"testing"
	"time"
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
			go client.StartConn()
			var exit chan string
			time.AfterFunc(5*time.Second, func() {
				var data []LoadBaseFactorBody
				_ = client.LoadBaseFactor(data)
				time.Sleep(5 * time.Second)
				close(exit)
			})
			<-exit
		})
	}
}
