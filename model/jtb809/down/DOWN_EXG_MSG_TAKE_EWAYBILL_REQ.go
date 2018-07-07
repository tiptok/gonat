package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_EXG_MSG_TAKE_EWAYBILL_REQ 上报车辆电子运单请求  920B
type DOWN_EXG_MSG_TAKE_EWAYBILL_REQ struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
}

func (e *DOWN_EXG_MSG_TAKE_EWAYBILL_REQ) GetMsgId() interface{} {
	return model.J从链路动态信息交换
}
func (e *DOWN_EXG_MSG_TAKE_EWAYBILL_REQ) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *DOWN_EXG_MSG_TAKE_EWAYBILL_REQ) GetDBSql() string {
	return ""
}
