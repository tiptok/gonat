package up

import (
	"time"

	"github.com/tiptok/gonat/model"
)

//UP_EXG_MSG_APPLY_HISGNSSDATA_REQ 补发车辆定位信息请求 1209
type UP_EXG_MSG_APPLY_HISGNSSDATA_REQ struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
	Start_Time    time.Time
	End_Time      time.Time
}
