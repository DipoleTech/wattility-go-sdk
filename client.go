package wattility_go_sdk

import (
	"github.com/go-resty/resty/v2"
	"strconv"
	"time"
)

type Client struct {
	Client    *resty.Client
	Host      string
	AppId     string
	AppSecret string
	sign      *Sign
}

func NewClient(appId, appSecret string) (*Client, error) {
	if len(appId) != 16 || len(appSecret) != 32 {
		return nil, nil
	}
	sign := NewSign(appSecret)
	client := &Client{
		AppId:     appId,
		AppSecret: appSecret,
		sign:      sign,
		Client:    resty.New(),
		Host:      "http://localhost:9001",
	}

	return client, nil
}

func (c *Client) do(uri, method string, body []byte) (res *resty.Response, err error) {
	ts := time.Now().UTC().Unix()
	random := "123"
	content := c.sign.Content(method, uri, random, string(body), ts)
	sign, _ := c.sign.Encrypt([]byte(content))
	signBase64 := c.sign.SignBase64(sign)
	signHash := c.sign.SignHash(signBase64)
	md5Str := c.sign.ContentMD5([]byte(content))

	c.Client.SetHeaders(map[string]string{
		"x-oauth-app-id":    c.AppId,
		"x-oauth-random":    random,
		"x-oauth-ts":        strconv.FormatInt(ts, 10),
		"x-oauth-md5":       md5Str,
		"x-oauth-sign-hash": signHash,
		"Content-Type":      "application/json",
	})

	res, err = c.Client.R().SetBody(body).Execute(method, c.Host+uri)
	return
}
