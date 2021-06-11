package wattility_go_sdk

func (c *Client) AppSHInfoSync(body AppSHInfoSyncBody) error {
	return c.do(2000, body)
}
func (c *Client) MomentDataReport(body MomentDataSyncBody) error {
	return c.do(2003, body)
}
