package api

type DrResponse struct {
	Root        string `json:"root"`
	Version     int    `json:"version"`
	Code        int    `json:"code"`
	Description string `json:"description"`
	RequestID   string `json:"requestID"`
	DnID        string `json:"dnID"`
}

type CreateRegistrationResponse struct {
	DrResponse
	UnID            string        `json:"unID"`
	Transport       []string      `json:"transport"`
	ServiceSpecific []interface{} `json:"serviceSpecific"`
}

type RegisterReportResponse struct {
	DrResponse
	ReportRequest []ReportRequest `json:"reportRequest"`
}

type ResourceReportResponse struct {
	DrResponse
}

type QueryEventResponse struct {
	DrResponse
	Events []Event `json:"events"`
}

type DataReportResponse struct {
	DrResponse
}
