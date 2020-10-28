package wattility_go_sdk

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

type Sign struct {
	secret   string
	keyBlock cipher.Block
	iv       string
}

func NewSign(secret string) *Sign {
	size := aes.BlockSize
	keyBlock, _ := aes.NewCipher([]byte(secret))
	return &Sign{
		secret:   secret,
		keyBlock: keyBlock,
		iv:       secret[:size],
	}
}

func (sa *Sign) Content(method, rawQuery, random, body string, ts int64) string {
	var content = fmt.Sprint(
		method, "\n",
		rawQuery, "\n",
		random, "\n",
		ts, "\n",
		body,
	)
	return content
}

func (sa *Sign) Encrypt(content []byte) ([]byte, error) {
	paddinglen := aes.BlockSize - (len(content) % aes.BlockSize)
	for i := 0; i < paddinglen; i++ {
		content = append(content, byte(paddinglen))
	}
	enbuf := make([]byte, len(content))
	cbce := cipher.NewCBCEncrypter(sa.keyBlock, []byte(sa.iv))
	cbce.CryptBlocks(enbuf, content)
	return enbuf, nil
}

func (sa *Sign) SignBase64(sign []byte) string {
	return base64.StdEncoding.EncodeToString(sign)
}

func (sa *Sign) SignHash(signBase64 string) string {
	h := sha256.New()
	h.Write([]byte(signBase64))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func (sa *Sign) ContentMD5(content []byte) string {
	has := md5.Sum(content)
	return fmt.Sprintf("%x", has)
}
