package wattility_go_sdk

import "github.com/DipoleTech/wattility-go-sdk/sh"

func (c *Client) MetaDataReport(body sh.CreateRegistrationRequest) error {
	return c.do(2001, body)
}
func (c *Client) ResourceReport(body sh.ResourceReportRequest) error {
	return c.do(2002, body)
}
func (c *Client) MomentDataReport(body sh.MomentDataReportRequest) error {
	return c.do(2003, body)
}
