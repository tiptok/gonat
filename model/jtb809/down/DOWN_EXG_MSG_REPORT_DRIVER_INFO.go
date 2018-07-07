package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_EXG_MSG_REPORT_DRIVER_INFO 上报车辆驾驶员身份识别信息请求  920A
type DOWN_EXG_MSG_REPORT_DRIVER_INFO struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
}

func (e *DOWN_EXG_MSG_REPORT_DRIVER_INFO) GetMsgId() interface{} {
	return model.J从链路动态信息交换
}
func (e *DOWN_EXG_MSG_REPORT_DRIVER_INFO) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *DOWN_EXG_MSG_REPORT_DRIVER_INFO) GetDBSql() string {
	return ""
}
