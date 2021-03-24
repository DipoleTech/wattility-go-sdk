package wattility_go_sdk

import "time"

type AuthBody struct {
	AppId     string `json:"app_id"`
	Timestamp int64  `json:"timestamp"`
}

type LoadBasePredictionDaily struct {
	HouseholdNumber string `json:"household_number"` // 户号
	Prediction      []struct {
		RecordAt       time.Time `json:"record_at"`       // 当前功率时间点
		PredictedValue float64   `json:"predicted_value"` // 瞬时功率 预测值
	} `json:"prediction"`
}

type LoadBasePredictionRes struct {
	OrderId       uint64
	HouseholdData []struct {
		HouseholdNumber string `json:"household_number"` // 户号
		Prediction      []struct {
			RecordAt       time.Time `json:"record_at"`       // 当前功率时间点
			PredictedValue float64   `json:"predicted_value"` // 瞬时功率 预测值
		} `json:"prediction"`
	}
}
type LoadBasePredictionReq struct {
	OrderId         uint64
	Start           time.Time
	End             time.Time
	HouseholdNumber []string `json:"household_number"`
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

// 指令创建后服务端发送指令
type OrderPendingBody struct {
	OrderId       uint64    `json:"order_id"`
	OrderName     string    `json:"order_name"`
	OrderType     string    `json:"order_type"`
	ResponseType  string    `json:"response_type"`
	TargetPower   float64   `json:"target_power"` // 响应目标量
	StartAt       time.Time `json:"start_at"`     // 响应开始时间
	ExpiredAt     time.Time `json:"expired_at"`   // 响应结束时间
	DeadlineAt    time.Time `json:"deadline_at"`  // 截止时间
	HouseholdData []struct {
		HouseholdNumber string  `json:"household_number"`
		Power           float64 `json:"power"` // 分配量
		Summary         []struct {
			RecordAt       string  `json:"record_at"`       // 当前功率时间点
			PredictedValue float64 `json:"predicted_value"` // 瞬时功率 预测值
			BaseValue      float64 `json:"base_value"`      // 基线值
		}
	} `json:"household_data"`
}

// 指令创建后客户端发送策略
type OrderPendingStrategyBody struct {
	OrderId       uint64 `json:"order_id"`
	HouseholdData []struct {
		HouseholdNumber string `json:"household_number"` // 户号
		Strategy        string `json:"strategy"`         // 策略
	} `json:"household_data"`
}

// 指令确认发布后服务端发送策略
type OrderConfirmedBody struct {
	OrderId       uint64 `json:"order_id"`
	HouseholdData []struct {
		HouseholdNumber string `json:"household_number"` // 户号
		Strategy        string `json:"strategy"`         // 策略
		CustomStrategy  string `json:"custom_strategy"`  // 自定义策
	} `json:"household_data"`
}

// 指令结束后服务端发送数据
type OrderSettleBody struct {
	OrderId       uint64                     `json:"order_id"`
	HouseholdData []OrderSettleBodyHousehold `json:"household_data"`
}

type OrderSettleBodyHousehold struct {
	HouseholdNumber string                    `json:"household_number"` // 户号
	CorrectPercent  int                       `json:"correct_percent"`  // 准确率
	Summary         []OrderSettleBodySummary  `json:"summary"`
	Statistics      OrderSettleBodyStatistics `json:"statistics"`
}

type OrderSettleBodySummary struct {
	RecordAt     time.Time `json:"record_at"`     // 当前功率时间点
	PredictValue float64   `json:"predict_value"` // 瞬时功率 预测值
	BaseValue    float64   `json:"base_value"`    // 基线值
	RealValue    float64   `json:"real_value"`    // 实时值
	SettleValue  float64   `json:"settle_value"`  // 结算值
}

type OrderSettleBodyStatistics struct {
	SumPower    float64 `json:"sum_power"`    // 累计响应量
	AvgPower    float64 `json:"avg_power"`    // 平均响应量
	TargetPower float64 `json:"target_power"` // 目标响应量
}
