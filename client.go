package wattility_go_sdk

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"math/rand"
	"strconv"
	"time"
)

type Client struct {
	client    *resty.Client
	host      string
	appId     string
	appSecret string
	sign      *Sign
}

var (
	ErrorApp = errors.New("app id or app secret is error")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

type ResBody struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewClient(appId, appSecret string) (*Client, error) {
	if len(appId) != 16 || len(appSecret) != 32 {
		return nil, ErrorApp
	}
	sign, _ := NewSign(appSecret)
	c := &Client{
		appId:     appId,
		appSecret: appSecret,
		sign:      sign,
		client:    resty.New(),
		host:      "http://localhost:9001",
	}
	c.client.OnAfterResponse(func(client *resty.Client, response *resty.Response) error {
		var data ResBody
		err := json.Unmarshal(response.Body(), &data)
		if err != nil {
			return err
		}
		if data.Code != 0 {
			return errors.New(data.Message)
		}
		return nil
	})

	return c, nil
}

func (c *Client) SetDevHost(host string) {
	c.host = host
}

func (c *Client) do(uri, method string, body []byte) (res *resty.Response, err error) {
	ts := time.Now().UTC().Unix()
	random := RandString()
	content := c.sign.Content(method, uri, random, string(body), ts)
	sign := c.sign.Encrypt([]byte(content))
	signBase64 := c.sign.SignBase64(sign)
	signHash := c.sign.SignHash(signBase64)
	md5Str := c.sign.ContentMD5([]byte(content))

	c.client.SetHeaders(map[string]string{
		"x-oauth-app-id":    c.appId,
		"x-oauth-random":    random,
		"x-oauth-ts":        strconv.FormatInt(ts, 10),
		"x-oauth-md5":       md5Str,
		"x-oauth-sign-hash": signHash,
		"Content-Type":      "application/json",
	})

	res, err = c.client.R().SetBody(body).Execute(method, c.host+uri)
	return
}

func RandString() string {
	var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 10)
	for i := range b {
		b[i] = letters[rand.Intn(10)]
	}
	return string(b)
}
