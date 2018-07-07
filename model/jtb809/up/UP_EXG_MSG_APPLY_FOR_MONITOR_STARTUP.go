package up

import (
	"time"

	"github.com/tiptok/gonat/model"
)

//UP_EXG_MSG_APPLY_FOR_MONITOR_STARTUP UP_EXG_MSG_APPLY_FOR_MONITOR_STARTUP 1207
type UP_EXG_MSG_APPLY_FOR_MONITOR_STARTUP struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
	Start_Time    time.Time
	End_Time      time.Time
}
