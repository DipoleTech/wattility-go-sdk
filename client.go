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
	recHandle  *ReceiveHandle
}

type ReceiveHandle struct {
	LoadBasePredictionRec   func(receive []byte)
	LoadBaseFactorRec       func(receive []byte)
	OrderInfoRec            func(receive []byte)
	OrderCustomeStrategyRec func(receive []byte)
	OrderFinishRec          func(receive []byte)
}

var (
	ErrorApp = errors.New("app id or app secret is error")
)

//func init() {
//	rand.Seed(time.Now().UnixNano())
//}

func NewClient(appId, appSecret string) (*Client, error) {
	if len(appId) != 16 || len(appSecret) != 32 {
		return nil, ErrorApp
	}
	host := "localhost:8999"
	sign, _ := NewSign(appSecret)
	c := &Client{
		appId:      appId,
		appSecret:  appSecret,
		sign:       sign,
		socketConn: nil,
		host:       host,
		logger:     NewLogger(false),
		recHandle:  nil,
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
	c.recHandle = handle
}

func (c *Client) StartConn() {
	conn, err := net.Dial("tcp", c.host)
	if err != nil {
		c.logger.Print(err.Error())
		return
	}
	c.socketConn = conn
	for {
		dp := NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.socketConn, headData)
		if err != nil {
			c.logger.Print("read head error")
			break
		}

		//将headData字节流 拆包到msg中
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			c.logger.Print(fmt.Sprint("server unpack err:", err))
			return
		}
		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead
			msg.Data = make([]byte, msg.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(c.socketConn, msg.Data)
			if err != nil {
				c.logger.Print(fmt.Sprint("server unpack data err:", err))
				return
			}
			c.logger.Print(fmt.Sprint("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data)))
			switch msg.Id {
			case 0:
				c.Auth(string(msg.Data))
			case 1001:
				if c.recHandle.LoadBasePredictionRec != nil {
					go c.recHandle.LoadBasePredictionRec(msg.Data)
				}
			case 1002:
				if c.recHandle.LoadBaseFactorRec != nil {
					go c.recHandle.LoadBaseFactorRec(msg.Data)
				}
			case 1004:
				if c.recHandle.OrderInfoRec != nil {
					go c.recHandle.OrderInfoRec(msg.Data)
				}
			case 1005:
				if c.recHandle.OrderCustomeStrategyRec != nil {
					go c.recHandle.OrderCustomeStrategyRec(msg.Data)
				}
			case 1006:
				if c.recHandle.OrderFinishRec != nil {
					go c.recHandle.OrderFinishRec(msg.Data)
				}
			default:
			}
		}
	}
}
