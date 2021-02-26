package wattility_go_sdk

type AuthBody struct {
	AppId     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
}

type LoadBaseSummaryBody struct {
	HouseholdNumber string  `json:"household_number"` // 户号
	DatetimeStart   string  `json:"datetime_start"`   // 当前功率时间点
	DatetimeEnd     string  `json:"datetime_end"`     // 当前功率时间点
	ValueBase       float64 `json:"value_base"`       // 功率-工作日-平均值
	ValueReal       float64 `json:"value_real"`       // 功率-工作日-最近一次数据
	ValueWBase      float64 `json:"value_w_base"`     // 功率-周末-平均值
	ValueWReal      float64 `json:"value_w_real"`     // 功率-周末-最近一次数据
	ValueHBase      float64 `json:"value_h_base"`     // 功率-节假日-平均值
	ValueHReal      float64 `json:"value_h_real"`     // 功率-节假日-最近一次数据
}

type LoadBaseFactorBody struct {
	HouseholdNumber   string  `json:"household_number"`     // 户号
	DeviceSn          string  `json:"device_sn"`            // 设备号
	DeviceName        string  `json:"device_name"`          // 设备名
	Key               string  `json:"key"`                  // 属性key
	AttrName          string  `json:"attr_name"`            // 属性名
	AdjustableTime    int     `json:"adjustable_time"`      // 此设备响应时间，单位秒
	AdjustableComplex int     `json:"adjustable_complex"`   // 此设备响应难度，分1,2,3级
	SpringUpMaxLoad   float64 `json:"spring_up_max_load"`   // 春季
	SpringUpMinLoad   float64 `json:"spring_up_min_load"`   // 春季
	SpringDownMaxLoad float64 `json:"spring_down_max_load"` // 春季
	SpringDownMinLoad float64 `json:"spring_down_min_load"` // 春季
	SummerUpMaxLoad   float64 `json:"summer_up_max_load"`   // 夏季
	SummerUpMinLoad   float64 `json:"summer_up_min_load"`   // 夏季
	SummerDownMaxLoad float64 `json:"summer_down_max_load"` // 夏季
	SummerDownMinLoad float64 `json:"summer_down_min_load"` // 夏季
	AutumnUpMaxLoad   float64 `json:"autumn_up_max_load"`   // 秋季
	AutumnUpMinLoad   float64 `json:"autumn_up_min_load"`   // 秋季
	AutumnDownMaxLoad float64 `json:"autumn_down_max_load"` // 秋季
	AutumnDownMinLoad float64 `json:"autumn_down_min_load"` // 秋季
	WinterUpMaxLoad   float64 `json:"winter_up_max_load"`   // 冬季
	WinterUpMinLoad   float64 `json:"winter_up_min_load"`   // 冬季
	WinterDownMaxLoad float64 `json:"winter_down_max_load"` // 冬季
	WinterDownMinLoad float64 `json:"winter_down_min_load"` // 冬季
	AdviceSpringDown  string  `json:"advice_spring_down"`   // 春季调低操作建议，例：风速降低2挡,温度降低2°
	AdviceSpringUp    string  `json:"advice_spring_up"`
	AdviceSummerDown  string  `json:"advice_summer_down"`
	AdviceSummerUp    string  `json:"advice_summer_up"`
	AdviceAutumnDown  string  `json:"advice_autumn_down"`
	AdviceAutumnUp    string  `json:"advice_autumn_up"`
	AdviceWinterDown  string  `json:"advice_winter_down"`
	AdviceWinterUp    string  `json:"advice_winter_up"`
}
