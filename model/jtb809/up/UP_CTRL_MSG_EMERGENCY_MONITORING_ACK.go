package up

import (
	"github.com/tiptok/gonat/model"
)
//UP_CTRL_MSG_EMERGENCY_MONITORING_ACK 车辆应急接入督管平台应答 1505
type UP_CTRL_MSG_EMERGENCY_MONITORING_ACK struct{
    model.EntityBase
    Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
    RESULT byte
}