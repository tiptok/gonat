package up

import (
	"github.com/tiptok/gonat/model"
)

//UP_CTRL_MSG_EMERGENCY_MONITORING_ACK 车辆应急接入督管平台应答 1505
type UP_CTRL_MSG_EMERGENCY_MONITORING_ACK struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
	RESULT        byte
}

func (e *UP_CTRL_MSG_EMERGENCY_MONITORING_ACK) GetMsgId() interface{} {
	return model.J车辆单向监听应答
}
func (e *UP_CTRL_MSG_EMERGENCY_MONITORING_ACK) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *UP_CTRL_MSG_EMERGENCY_MONITORING_ACK) GetDBSql() string {
	return ""
}
