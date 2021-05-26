package wattility_go_sdk

func (c *Client) LoadBasePrediction(body LoadBasePredictionRes) error {
	return c.do(1001, body)
}

func (c *Client) LoadBaseFactor(body []LoadBaseFactorBody) error {
	return c.do(1002, body)
}

func (c *Client) LoadBasePredictionDaily(body []LoadBasePredictionDaily) error {
	return c.do(1003, body)
}

func (c *Client) OrderStrategy(body OrderStrategyBody) error {
	return c.do(1005, body)
}

func (c *Client) AppInfoSync(body AppSyncBody) error {
	return c.do(1007, body)
}
