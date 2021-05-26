package wattility_go_sdk

import "github.com/DipoleTech/wattility-go-sdk/sh"

func (c *Client) MomentDataReport(body sh.MomentDataReportRequest) error {
	return c.do(2001, body)
}

func (c *Client) IntervalDataReport(body sh.IntervalDataReportRequest) error {
	return c.do(2002, body)
}

func (c *Client) LoadForecastReport(body sh.LoadForecastReportRequest) error {
	return c.do(2003, body)
}

func (c *Client) PowerForecastReport(body sh.PowerForecastReportRequest) error {
	return c.do(2004, body)
}
