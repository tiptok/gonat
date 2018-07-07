package up

import (
	"github.com/tiptok/gonat/model"
)

//UP_EXG_MSG_APPLY_FOR_MONITOR_END 取消交换指定车辆定位信息请求 1208
type UP_EXG_MSG_APPLY_FOR_MONITOR_END struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
}
