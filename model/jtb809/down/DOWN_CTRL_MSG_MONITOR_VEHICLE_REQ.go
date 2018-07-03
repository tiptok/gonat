package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_CTRL_MSG_MONITOR_VEHICLE_REQ 车辆单向监听请求 9501
type DOWN_CTRL_MSG_MONITOR_VEHICLE_REQ struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
	MONITOR_TEL   string //回拨电话号码
}

func (e *DOWN_CTRL_MSG_MONITOR_VEHICLE_REQ) GetMsgId() interface{} {
	return model.J从链路车辆监管
}
func (e *DOWN_CTRL_MSG_MONITOR_VEHICLE_REQ) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *DOWN_CTRL_MSG_MONITOR_VEHICLE_REQ) GetDBSql() string {
	return ""
}
