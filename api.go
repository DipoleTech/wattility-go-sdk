package wattility_go_sdk

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
)

func (c *Client) CheckSignApi() (*resty.Response, error) {
	var url = "/v1/open_api/check_access"
	res, err := c.do(url, "GET", []byte(""))
	if err != nil {
		return res, err
	}
	return res, nil
}

func (c *Client) LoadBaseSummary(body LoadBaseSummaryBody) error {
	var url = "/v1/open_api/load_base/summary_chart"
	b, _ := json.Marshal(body)

	_, err := c.do(url, "POST", b)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) LoadBaseInfo(body LoadBaseInfoyBody) error {
	var url = "/v1/open_api/load_base/info"
	b, _ := json.Marshal(body)

	_, err := c.do(url, "POST", b)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) LoadBaseFactor(body []LoadBaseFactorBody) error {
	var url = "/v1/open_api/load_base/factor"
	b, _ := json.Marshal(body)

	_, err := c.do(url, "POST", b)
	if err != nil {
		return err
	}
	return nil
}
