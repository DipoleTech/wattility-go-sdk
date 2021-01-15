package wattility_go_sdk

import (
	"encoding/json"
	"github.com/aceld/zinx/znet"
	"time"
)

type APIData struct {
	AppID       string `json:"app_id"`
	SignContent []byte `json:"sign_content"`
}

func (c *Client) Auth(messageData string) {
	switch messageData {
	case "auth":
		dp := znet.NewDataPack()

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

		msg, _ := dp.Pack(znet.NewMsgPackage(0, msgData))
		_, _ = c.socketConn.Write(msg)
	case "OK":
		return
	}

}

func (c *Client) LoadBaseSummary(body []LoadBaseSummaryBody) error {
	b, _ := json.Marshal(body)
	bs := c.sign.Encrypt(b)
	dp := znet.NewDataPack()

	var data = APIData{
		AppID:       c.appId,
		SignContent: bs,
	}

	msgData, _ := json.Marshal(data)

	msg, _ := dp.Pack(znet.NewMsgPackage(0, msgData))
	_, _ = c.socketConn.Write(msg)
	return nil
}

func (c *Client) LoadBaseFactor(body []LoadBaseFactorBody) error {
	b, _ := json.Marshal(body)
	bs := c.sign.Encrypt(b)
	dp := znet.NewDataPack()

	var data = APIData{
		AppID:       c.appId,
		SignContent: bs,
	}

	msgData, _ := json.Marshal(data)

	msg, _ := dp.Pack(znet.NewMsgPackage(0, msgData))
	_, _ = c.socketConn.Write(msg)
	return nil
}
