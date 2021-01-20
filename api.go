package wattility_go_sdk

import (
	"encoding/json"
	"time"
)

type APIData struct {
	AppID       string `json:"app_id"`
	SignContent []byte `json:"sign_content"`
}

func (c *Client) Auth(messageData string) {
	switch messageData {
	case "auth":
		dp := NewDataPack()

		var body = AuthBody{
			AppId:     c.appId,
			Timestamp: time.Now().Unix(),
		}
		b, _ := json.Marshal(body)
		bs := c.sign.Encrypt(b)

		var data = APIData{
			AppID:       c.appId,
			SignContent: bs,
		}
		msgData, _ := json.Marshal(data)

		msg, _ := dp.Pack(NewMsgPackage(0, msgData))
		_, _ = c.socketConn.Write(msg)
	case "OK":
		return
	}

}

func (c *Client) do(id uint32, body interface{}) error {
	b, err := json.Marshal(body)
	if err != nil {
		return err
	}

	bs := c.sign.Encrypt(b)
	dp := NewDataPack()

	var data = APIData{
		AppID:       c.appId,
		SignContent: bs,
	}

	msgData, err := json.Marshal(data)
	if err != nil {
		return err
	}

	msg, err := dp.Pack(NewMsgPackage(id, msgData))
	if err != nil {
		return err
	}

	_, err = c.socketConn.Write(msg)
	return err
}

func (c *Client) LoadBaseSummary(body []LoadBaseSummaryBody) error {
	return c.do(1001, body)
}

func (c *Client) LoadBaseFactor(body []LoadBaseFactorBody) error {
	return c.do(1002, body)
}
