package wattility_go_sdk

func (c *Client) CheckAccessAPI() error {
	var url = "/v1/open_api/check_access"
	_, err := c.do(url, "GET", nil)
	if err != nil {
		return err
	}
	return nil
}
