package sh

type DrRequest struct {
	Root      string `json:"root"`
	Version   int    `json:"version"`
	RequestID string `json:"requestID"`
	DnID      string `json:"dnID"`
}

type CreateRegistrationRequest struct {
	DrRequest
	RegistrationID   string `json:"registrationID"`
	DnName           string `json:"dnName"`
	ReportOnly       bool   `json:"reportOnly"`
	PullMode         bool   `json:"pullMode"`
	Signature        bool   `json:"signature"`
	Transport        string `json:"transport"`
	TransportAddress string `json:"transportAddress"`
}

type QueryRegistrationRequest struct {
	DrRequest
}

type RegisterReportRequest struct {
	DrRequest
	ReportRequestID string   `json:"reportRequestID"`
	Report          []Report `json:"report"`
}

type ResourceReportRequest struct {
	DrRequest
	ReportRequestID string     `json:"reportRequestID"`
	Resource        []Resource `json:"resource"`
}

type QueryEventRequest struct {
	DrRequest
	ReplyLimit int    `json:"replyLimit"`
	Dtstart    string `json:"dtstart"`
	Duration   string `json:"duration"`
}

type MomentDataReportRequest struct {
	DrRequest
	ReportRequestID string      `json:"reportRequestID"`
	CreatedDateTime string      `json:"createdDateTime"`
	PointData       []PointData `json:"pointData"`
}

type IntervalDataReportRequest struct {
	DrRequest
	ReportRequestID string           `json:"reportRequestID"`
	CreatedDateTime string           `json:"createdDateTime"`
	PointCurveData  []PointCurveData `json:"pointCurveData"`
}

type LoadForecastReportRequest struct {
	DrRequest
	ReportRequestID   string              `json:"reportRequestID"`
	CreatedDateTime   string              `json:"createdDateTime"`
	ResourceCurveData []ResourceCurveData `json:"resourceCurveData"`
}

type PowerForecastReportRequest struct {
	DrRequest
	ReportRequestID   string              `json:"reportRequestID"`
	CreatedDateTime   string              `json:"createdDateTime"`
	ResourceCurveData []ResourceCurveData `json:"resourceCurveData"`
}

type CreateOptRequest struct {
	DrRequest
	OptID           string  `json:"optID"`
	OptType         OptType `json:"optType"`
	OptReason       string  `json:"optReason"`
	MarketContext   string  `json:"marketContext"`
	CreatedDateTime string  `json:"createdDateTime"`
}

type CancelOptRequest struct {
	DrRequest
	OptID string `json:"optID"`
}
