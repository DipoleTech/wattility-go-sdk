package wattility_go_sdk

import "time"

type LoadBaseSummaryBody struct {
	RecordAt      time.Time `json:"record_time"`
	BaseLineValue float64   `json:"base_line_value"`
	RealTimeValue float64   `json:"real_time_value"`
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
	Name string `json:"name"`
	Tag  string `json:"tag"`
	LoadBaseInfoyBody
}
