package up

import (
	"github.com/tiptok/gonat/model"
)
//UP_CTRL_MSG_TAKE_TRAVEL_ACK 上报车辆行驶记录应答 1504
type UP_CTRL_MSG_TAKE_TRAVEL_ACK struct{
    model.EntityBase
    Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
    TRAVELDATA_INFO string //车辆行驶记录信息
    COMMAND_TYPE byte//命令字
}