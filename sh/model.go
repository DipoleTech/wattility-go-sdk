package api

type Properties struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Descriptor struct {
	EventID              string `json:"eventID"`
	Status               string `json:"status"`
	CreatedDateTime      string `json:"createdDateTime"`
	ModificationDateTime string `json:"modificationDateTime"`
	ModificationNumber   int    `json:"modificationNumber"`
	ModificationReason   string `json:"modificationReason"`
	Priority             int    `json:"priority"`
	Test                 bool   `json:"test"`
	Comment              string `json:"comment"`
	MarketURL            string `json:"marketURL"`
}

type ActivePeriod struct {
	Dtstart      string `json:"dtstart"`
	Duration     string `json:"duration"`
	Notification string `json:"notification"`
	Rampup       string `json:"rampup"`
	Recovery     string `json:"recovery"`
}

type Intervals struct {
	Irregular struct {
		Values []struct {
			Value     float64 `json:"value"`
			Timestamp string  `json:"timestamp"`
		} `json:"values"`
	} `json:"irregular"`
}

type EventBaseline struct {
	BaselineID   string    `json:"baselineID"`
	BaselineName string    `json:"baselineName"`
	Metric       Metric    `json:"metric"`
	Intervals    CurveData `json:"intervals"`
}

type EventSignal struct {
	SignalID   string    `json:"signalID"`
	SignalName string    `json:"signalName"`
	SignalType string    `json:"signalType"`
	Metric     Metric    `json:"metric"`
	Intervals  CurveData `json:"intervals"`
}

type Signals struct {
	Baseline EventBaseline `json:"baseline"`
	Signal   []EventSignal `json:"signal"`
}

type Event []struct {
	ResponseRequired string       `json:"responseRequired"`
	Descriptor       Descriptor   `json:"descriptor"`
	ActivePeriod     ActivePeriod `json:"activePeriod"`
	Signals          Signals      `json:"signals"`
	Target           interface{}  `json:"target"`
}

type Metric struct {
	MetricName string `json:"metricName"`
	Multiplier string `json:"multiplier"`
	Symbol     string `json:"symbol"`
}

type EndDeviceAsset struct {
	Mrid string `json:"mrid"`
}

type ReportSubject struct {
	EndDeviceAsset EndDeviceAsset `json:"endDeviceAsset"`
}

type ReportDataSource struct {
	ResourceID []string `json:"resourceID"`
}

type SamplingRate struct {
	OnChange  bool   `json:"onChange"`
	MinPeriod string `json:"minPeriod"`
	MaxPeriod string `json:"maxPeriod"`
}

type ReportDescription struct {
	RID              int              `json:"rID"`
	ReadingType      string           `json:"readingType"`
	Metric           Metric           `json:"metric"`
	ReportSubject    ReportSubject    `json:"reportSubject"`
	ReportDataSource ReportDataSource `json:"reportDataSource"`
	SamplingRate     SamplingRate     `json:"samplingRate"`
}

type Report struct {
	CreatedDateTime   string            `json:"createdDateTime"`
	ReportDescription ReportDescription `json:"reportDescription"`
}

type ReportSpecifier struct {
	ReportType   string `json:"reportType"`
	BackDuration string `json:"backDuration"`
	Period       string `json:"period"`
	Points       []int  `json:"points"`
}

type ReportRequest []struct {
	ReportRequestID string          `json:"reportRequestID"`
	ReportSpecifier ReportSpecifier `json:"reportSpecifier"`
}

type Resource struct {
	ResourceID       string         `json:"resourceID"`
	ResourceName     string         `json:"resourceName"`
	Status           string         `json:"status"`
	ParentResourceID string         `json:"parentResourceID"`
	EndDeviceAsset   EndDeviceAsset `json:"endDeviceAsset"`
	Properties       []Properties   `json:"properties"`
}

type PointData struct {
	RID       int     `json:"rID"`
	Value     float64 `json:"value"`
	Timestamp string  `json:"timestamp"`
	Quality   string  `json:"quality"`
}

type Data struct {
	Value     float64 `json:"value"`
	Timestamp string  `json:"timestamp"`
	Quality   string  `json:"quality"`
}

type Irregular struct {
	Values []Data `json:"values"`
}

type PointCurveData struct {
	RID       int       `json:"rID"`
	Irregular Irregular `json:"irregular"`
}

type RegularCurve struct {
	Dtstart string    `json:"dtstart"`
	Period  string    `json:"period"`
	Array   []float64 `json:"array"`
}

type IrregularCurve struct {
	Values []Data `json:"values"`
}

type CurveData struct {
	Irregular IrregularCurve `json:"irregular"`
	Regular   RegularCurve   `json:"regular"`
}

type ResourceCurveData struct {
	ResourceID string       `json:"resourceID"`
	Regular    RegularCurve `json:"regular"`
}
