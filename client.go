package wattility_go_sdk

import (
	"errors"
	"fmt"
	"io"
	"net"
)

type Client struct {
	appId      string
	appSecret  string
	host       string
	socketConn net.Conn
	logger     Logger
	sign       *Sign
	recRouter  map[uint32]func([]byte)
	DisConnect func()
}

type ReceiveHandle struct {
	LoadBasePredictionRec   func(receive []byte)
	LoadBaseFactorRec       func(receive []byte)
	OrderInfoRec            func(receive []byte)
	OrderCustomeStrategyRec func(receive []byte)
	OrderFinishRec          func(receive []byte)
	AppInfoSync             func(receive []byte)
	MetaDataReport          func(receive []byte)
	ResourceReport          func(receive []byte)
}

var (
	ErrorApp = errors.New("app id or app secret is error")
)

func NewClient(appId, appSecret string) (*Client, error) {
	if len(appId) != 16 || len(appSecret) != 32 {
		return nil, ErrorApp
	}
	sign, _ := NewSign(appSecret)
	c := &Client{
		appId:      appId,
		appSecret:  appSecret,
		sign:       sign,
		socketConn: nil,
		host:       "localhost:8999",
		logger:     NewLogger(false),
		recRouter:  make(map[uint32]func([]byte)),
	}
	return c, nil
}

func (c *Client) SetHost(host string) {
	c.host = host
}

func (c *Client) SetDebug() {
	c.logger = NewLogger(true)
}

func (c *Client) SetRecHandle(handle *ReceiveHandle) {
	c.recRouter[1001] = handle.LoadBasePredictionRec
	c.recRouter[1002] = handle.LoadBaseFactorRec
	c.recRouter[1004] = handle.OrderInfoRec
	c.recRouter[1005] = handle.OrderCustomeStrategyRec
	c.recRouter[1006] = handle.OrderFinishRec
	c.recRouter[1007] = handle.AppInfoSync
	c.recRouter[2001] = handle.MetaDataReport
	c.recRouter[2002] = handle.ResourceReport
}

func (c *Client) StartConn() {
	conn, err := net.Dial("tcp", c.host)
	if err != nil {
		c.logger.Print(err.Error())
		if c.DisConnect != nil {
			c.DisConnect()
		}
		return
	}
	c.socketConn = conn
	for {
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.socketConn, headData)
		if err != nil {
			if err == io.EOF {
				if c.DisConnect != nil {
					c.DisConnect()
				}
				c.logger.Print("conn disconnect")
			}
			c.logger.Print(fmt.Sprintf("read head error: %v", err))
			break
		}

		msgHead, err := dp.Unpack(headData)
		if err != nil {
			c.logger.Print(fmt.Sprint("server unpack err:", err))
			return
		}
		if msgHead.GetDataLen() > 0 {
			msg := msgHead
			msg.Data = make([]byte, msg.GetDataLen())

			_, err := io.ReadFull(c.socketConn, msg.Data)
			if err != nil {
				c.logger.Print(fmt.Sprint("server unpack data err:", err))
				return
			}
			c.logger.Print(fmt.Sprint("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data)))
			switch msg.Id {
			case 0:
				c.Auth(string(msg.Data))
			default:
				handler, ok := c.recRouter[msg.Id]
				if ok && handler != nil {
					go handler(msg.Data)
				}
			}
		}
	}
	return
}
