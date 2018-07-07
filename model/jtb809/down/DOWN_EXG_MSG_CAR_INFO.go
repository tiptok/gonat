package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_EXG_MSG_CAR_INFO 交换车辆静态信息 9204
type DOWN_EXG_MSG_CAR_INFO struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
    Car_Info string 
}

func (e *DOWN_EXG_MSG_CAR_INFO) GetMsgId() interface{} {
	return model.J从链路动态信息交换
}
func (e *DOWN_EXG_MSG_CAR_INFO) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *DOWN_EXG_MSG_CAR_INFO) GetDBSql() string {
	return ""
}
