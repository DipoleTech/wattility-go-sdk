package wattility_go_sdk

import "time"

type AuthBody struct {
	AppId     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
}

// 预测线
type LoadBaseSummaryBody struct {
	HouseholdNumber string  `json:"household_number"` // 户号
	DatetimeAt      string  `json:""`                 // 当前功率时间点
	Power           float64 `json:"power"`            // 瞬时功率
}

// 发电因子
type LoadBaseFactorBody struct {
	HouseholdNumber   string  `json:"household_number"`     // 户号
	DeviceSn          string  `json:"device_sn"`            // 设备号
	DeviceName        string  `json:"device_name"`          // 设备名
	Spec              string  `json:"spec"`                 // 规格KEY
	SpecName          string  `json:"spec_name"`            // 规格名
	SpringUpMaxLoad   float64 `json:"spring_up_max_load"`   // 春季调高最高负载
	SpringUpMinLoad   float64 `json:"spring_up_min_load"`   // 春季调高最低负载
	SpringDownMaxLoad float64 `json:"spring_down_max_load"` // 春季调低最高负载
	SpringDownMinLoad float64 `json:"spring_down_min_load"` // 春季调低最低负载
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
}

type OrderBody struct {
	OrderId       uint64    `json:"order_id"`
	OrderName     string    `json:"order_name"`
	OrderType     string    `json:"order_type"`
	ResponseType  string    `json:"response_type"`
	TargetPower   float64   `json:"target_power"` // 响应目标量
	StartAt       time.Time `json:"start_at"`     // 响应开始时间
	ExpiredAt     time.Time `json:"expired_at"`   // 响应结束时间
	DeadlineAt    time.Time `json:"deadline_at"`  // 截止时间
	HouseholdData []struct {
		HouseholdNumber string   `json:"household_number"`
		CorrectPercent  int      `json:"correct_percent"` // 准确率
		Summary         []string `json:"summary"`
		Strategy        []string `json:"strategy"`
	} `json:"household_data"`
}
