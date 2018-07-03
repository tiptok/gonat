package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_CTRL_MSG_TAKE_TRAVEL_REQ 上报车辆行驶记录请求 0x9504
type DOWN_CTRL_MSG_TAKE_TRAVEL_REQ struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
	COMMAND_TYPE  byte   //命令字
}

func (e *DOWN_CTRL_MSG_TAKE_TRAVEL_REQ) GetMsgId() interface{} {
	return model.J从链路车辆监管
}
func (e *DOWN_CTRL_MSG_TAKE_TRAVEL_REQ) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *DOWN_CTRL_MSG_TAKE_TRAVEL_REQ) GetDBSql() string {
	return ""
}