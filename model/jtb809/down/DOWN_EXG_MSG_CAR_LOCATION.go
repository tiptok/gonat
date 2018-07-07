package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_EXG_MSG_CAR_LOCATION 交换车辆定位信息 9202
type DOWN_EXG_MSG_CAR_LOCATION struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
    GNSS_DATA_LIST []model.LocationInfoItem 
}

func (e *DOWN_EXG_MSG_CAR_LOCATION) GetMsgId() interface{} {
	return model.J从链路动态信息交换
}
func (e *DOWN_EXG_MSG_CAR_LOCATION) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *DOWN_EXG_MSG_CAR_LOCATION) GetDBSql() string {
	return ""
}
