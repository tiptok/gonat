package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_CTRL_MSG_MONITOR_VEHICLE_REQ 车辆单向监听请求 9501
type DOWN_CTRL_MSG_MONITOR_VEHICLE_REQ struct {
	model.EntityBase
	MONITOR_TEL string //回拨电话号码
}
