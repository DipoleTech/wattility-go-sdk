package sh

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
	MetricName MetricName `json:"metricName"`
	Multiplier string     `json:"multiplier"`
	Symbol     string     `json:"symbol"`
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
	ReadingType      ReadingType      `json:"readingType"`
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

type MeterAsset struct {
	Mrid string `json:"mrid"`
}

type Target struct {
	GroupID        []string         `json:"groupID"`
	ResourceID     []string         `json:"resourceID"`
	DnID           []string         `json:"dnID"`
	PartyID        []string         `json:"partyID"`
	EndDeviceAsset []EndDeviceAsset `json:"endDeviceAsset"`
	MeterAsset     []MeterAsset     `json:"meterAsset"`
}

// OptType enum
// optIn	参与
// optOut	不参与
type OptType string

// MetricName enum
// AP	有功功率
// AP_E	有功电能
// AP_PE	发电量
type MetricName string

// ReadingType enum
// Direct_Read	直接读数(常用于电表)
// Net	净值
// Allocated	分配值
// Estimated	估计值
// Summed	求和（常用于表示总体数值）
// Derived	推断值
// Mean	平均值
// Peak	最高值
// Hybrid	混合值，针对需求响应聚合商下辖用户，可能存在不同用户读取类型不同的情况
// Contract	合同值
// Projected	预测值（常用于表示总体预测）
// x-RMS	均方根
// x-notApplicable	不适用
type ReadingType string

type DistributeEventRequest struct {
	DrRequest
	Events []Event `json:"events"`
}
