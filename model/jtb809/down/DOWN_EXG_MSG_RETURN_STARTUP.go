package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_EXG_MSG_RETURN_STARTUP 结束车辆定位信息交换请求  9205
type DOWN_EXG_MSG_RETURN_STARTUP struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
    Reason_Code byte
}

func (e *DOWN_EXG_MSG_RETURN_STARTUP) GetMsgId() interface{} {
	return model.J从链路动态信息交换
}
func (e *DOWN_EXG_MSG_RETURN_STARTUP) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *DOWN_EXG_MSG_RETURN_STARTUP) GetDBSql() string {
	return ""
}
