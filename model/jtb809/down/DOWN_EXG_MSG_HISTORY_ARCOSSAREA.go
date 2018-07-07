package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_EXG_MSG_HISTORY_ARCOSSAREA 补发车辆定位信息应答  9209
type DOWN_EXG_MSG_HISTORY_ARCOSSAREA struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
    Result byte
}

func (e *DOWN_EXG_MSG_HISTORY_ARCOSSAREA) GetMsgId() interface{} {
	return model.J从链路动态信息交换
}
func (e *DOWN_EXG_MSG_HISTORY_ARCOSSAREA) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *DOWN_EXG_MSG_HISTORY_ARCOSSAREA) GetDBSql() string {
	return ""
}
