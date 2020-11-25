package wattility_go_sdk

type LoadBaseSummaryBody struct {
	DatetimeStart string  `json:"datetime_start"` // 当前功率时间点
	DatetimeEnd   string  `json:"datetime_end"`   // 当前功率时间点
	ValueBase     float64 `json:"value_base"`     // 功率-工作日-平均值
	ValueReal     float64 `json:"value_real"`     // 功率-工作日-最近一次数据
	ValueWBase    float64 `json:"value_w_base"`   // 功率-周末-平均值
	ValueWReal    float64 `json:"value_w_realz"`  // 功率-周末-最近一次数据
	ValueHBase    float64 `json:"value_h_base"`   // 功率-节假日-平均值
	ValueHReal    float64 `json:"value_h_real"`   // 功率-节假日-最近一次数据
}

type LoadBaseInfoyBody struct {
	SpringUpMaxLoad   float64 `json:"spring_up_max_load"`
	SpringUpMinLoad   float64 `json:"spring_up_min_load"`
	SpringDownMaxLoad float64 `json:"spring_down_max_load"`
	SpringDownMinLoad float64 `json:"spring_down_min_load"`
	SummerUpMaxLoad   float64 `json:"summer_up_max_load"`
	SummerUpMinLoad   float64 `json:"summer_up_min_load"`
	SummerDownMaxLoad float64 `json:"summer_down_max_load"`
	SummerDownMinLoad float64 `json:"summer_down_min_load"`
	AutumnUpMaxLoad   float64 `json:"autumn_up_max_load"`
	AutumnUpMinLoad   float64 `json:"autumn_up_min_load"`
	AutumnDwonMaxLoad float64 `json:"autumn_dwon_max_load"`
	AutumnDownMinLoad float64 `json:"autumn_down_min_load"`
	WinterUpMaxLoad   float64 `json:"winter_up_max_load"`
	WinterUpMinLoad   float64 `json:"winter_up_min_load"`
	WinterDownMaxLoad float64 `json:"winter_down_max_load"`
	WinterDownMinLoad float64 `json:"winter_down_min_load"`
}

type LoadBaseFactorBody struct {
	DeviceSn          string `json:"device_sn"`          //设备号
	DeviceName        string `json:"device_name"`        //设备名
	Key               string `json:"key"`                //属性key
	AttrName          string `json:"attr_name"`          //属性名
	AdjustableTime    int    `json:"adjustable_time"`    //此设备响应时间，单位秒
	AdjustableComplex int    `json:"adjustable_complex"` //此设备响应难度，分1,2,3级
	LoadBaseInfoyBody
	AdviceSpringDown string `json:"advice_spring_down"` //春季调低操作建议，例：风速降低2挡,温度降低2°
	AdviceSpringUp   string `json:"advice_spring_up"`
	AdviceSummerDown string `json:"advice_summer_down"`
	AdviceSummerUp   string `json:"advice_summer_up"`
	AdviceAutumnDown string `json:"advice_autumn_down"`
	AdviceAutumnUp   string `json:"advice_autumn_up"`
	AdviceWinterDown string `json:"advice_winter_down"`
	AdviceWinterUp   string `json:"advice_winter_up"`
}
