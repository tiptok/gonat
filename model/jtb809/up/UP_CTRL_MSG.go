package up

import (
	"github.com/tiptok/gonat/model"
)

//UP_CTRL_MSG 主链路车辆监管
type UP_CTRL_MSG struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
}
