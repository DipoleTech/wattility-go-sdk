package wattility_go_sdk

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"time"
)

func (c *Client) CheckSignApi() (*resty.Response, error) {
	var url = "/v1/open_api/check_access"
	res, err := c.do(url, "GET", []byte(""))
	if err != nil {
		return res, err
	}
	return res, nil
}

type LoadBaseSummaryBody struct {
	RecordAt      time.Time `json:"record_time"`
	BaseLineValue float64   `json:"base_line_value"`
	RealTimeValue float64   `json:"real_time_value"`
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

type LoadBaseInfoyBody struct {
	SpringMaxLoad float64 `json:"spring_max_load"`
	SpringMinLoad float64 `json:"spring_min_load"`
	SummerMaxLoad float64 `json:"summer_max_load"`
	SummerMinLoad float64 `json:"summer_min_load"`
	AutumnMaxLoad float64 `json:"autumn_max_load"`
	AutumnMinLoad float64 `json:"autumn_min_load"`
	WinterMaxLoad float64 `json:"winter_max_load"`
	WinterMinLoad float64 `json:"winter_min_load"`
}

func (c *Client) LoadBaseInfo(body LoadBaseInfoyBody) error {
	var url = "/v1/open_api/load_base/"
	b, _ := json.Marshal(body)

	_, err := c.do(url, "POST", b)
	if err != nil {
		return err
	}
	return nil
}
