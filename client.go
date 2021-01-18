package wattility_go_sdk

import (
	"errors"
	"fmt"
	"github.com/aceld/zinx/znet"
	"io"
	"math/rand"
	"net"
	"time"
)

type Client struct {
	socketConn         net.Conn
	socketHost         string
	dev                bool
	host               string
	appId              string
	appSecret          string
	sign               *Sign
	LoadBaseSummaryRec func(receive string)
	LoadBaseFactorRec  func(receive string)
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
	host := "http://localhost:9001"
	sign, _ := NewSign(appSecret)
	conn, _ := net.Dial("tcp", host)
	c := &Client{
		appId:      appId,
		appSecret:  appSecret,
		dev:        false,
		sign:       sign,
		socketConn: conn,
		host:       host,
	}
	return c, nil
}

func (c *Client) SetHost(host string) {
	conn, _ := net.Dial("tcp", host)
	c.socketConn = conn
	c.host = host
}

func (c *Client) StartConn() {
	for {
		dp := znet.NewDataPack()
		headData := make([]byte, dp.GetHeadLen())
		_, err := io.ReadFull(c.socketConn, headData)
		if err != nil {
			fmt.Println("read head error")
			break
		}

		//将headData字节流 拆包到msg中
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}
		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(c.socketConn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}
			fmt.Println("==> Recv Msg: ID=", msg.Id, ", len=", msg.DataLen, ", data=", string(msg.Data))
			switch msg.Id {
			case 0:
				fmt.Println("auth:", msg.Id)
				c.Auth(string(msg.Data))
			case 1001:
				if c.LoadBaseSummaryRec != nil {
					c.LoadBaseSummaryRec(string(msg.Data))
				}
			case 1002:
				if c.LoadBaseFactorRec != nil {
					c.LoadBaseFactorRec(string(msg.Data))
				}
			default:
			}
		}
	}

}
