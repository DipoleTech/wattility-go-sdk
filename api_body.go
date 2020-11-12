package wattility_go_sdk

import "time"

type LoadBaseSummaryBody struct {
	RecordAt      time.Time `json:"record_time"`
	BaseLineValue float64   `json:"base_line_value"`
	RealTimeValue float64   `json:"real_time_value"`
}

type LoadBaseInfoyBody struct {
	SpringMaxLoad float64 `json:"spring_max_load"`
	SpringMinLoad float64 `json:"spring_min_load"`
	SummerMaxLoad float64 `json:"summer_max_load"`
	SummerMinLoad float64 `json:"summer_min_load"`
	AutumnMaxLoad float64 `json:"autumn_max_load"`
	AutumnMinLoad float64 `json:"autumn_min_load"`
	WinterMaxLoad float64 `json:"winter_max_load"`
	WinterMinLoad float64 `json:"winter_min_load"`
}

type LoadBaseFactorBody struct {
	Name     string  `json:"name"`
	Label    string  `json:"label"`
	MaxValue float64 `json:"max_value"`
	MinValue float64 `json:"min_value"`
}
