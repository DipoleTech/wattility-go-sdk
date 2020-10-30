package wattility_go_sdk

import "github.com/go-resty/resty/v2"

func (c *Client) CheckAccessAPI() (*resty.Response, error) {
	var url = "/v1/open_api/check_access"
	res, err := c.do(url, "GET", []byte(""))
	if err != nil {
		return res, err
	}
	return res, nil
}
