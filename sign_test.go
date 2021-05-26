package wattility_go_sdk

import (
	"testing"
)

var secret = "z48BT1uqa10HUQeNK87bfEjita1bnQ7i"

func TestSign_Encrypt(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		want    []byte
	}{
		{"1", []byte("dsa"), []byte("dsa")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sa, _ := NewSign(secret)
			got := sa.Encrypt(tt.content)
			t.Log(got)
		})
	}
}

func TestSign_Decrypter(t *testing.T) {
	tests := []struct {
		name    string
		content []byte
		es      string
		ds      string
		want    []byte
	}{
		{"1", []byte("dsa"), "z48BT1uqa10HUQeNK87bfEjita1bnQ7i", "z48BT1uqa10HUQeNK87bfEjita1bnQ7i", []byte("dsa")},
		{"err", []byte("dsa"), "z48BT1uqa10HUQeNK87bfEjita1bnQ7i", "z48BT1uqa10HUQeNK87bfEjita1bnQ88", []byte("dsa")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Run(tt.name, func(t *testing.T) {
				sa, _ := NewSign(tt.es)
				eGot := sa.Encrypt(tt.content)

				dSign, _ := NewSign(tt.ds)
				got := dSign.Decrypt(eGot)
				t.Log(string(got))
				t.Log(got)
			})
		})
	}
}
