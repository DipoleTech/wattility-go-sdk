package wattility_go_sdk

import (
	"encoding/json"
	"time"
)

type APIData struct {
	AppID       string `json:"app_id"`
	Timestamp   int64  `json:"timestamp"`
	SignContent []byte `json:"sign_content"`
}

type AuthBody struct {
	AppId     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
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
		Timestamp:   time.Now().Unix(),
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
			Timestamp:   body.Timestamp,
			SignContent: bs,
		}
		msgData, _ := json.Marshal(data)

		msg, _ := dp.Pack(NewMsgPackage(0, msgData))
		_, _ = c.socketConn.Write(msg)
	case "OK":
		return
	}

}
