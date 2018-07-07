package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_EXG_MSG_APPLY_IIISGNSSDATA_ACK 车辆定位信息交换补发 9203
type DOWN_EXG_MSG_APPLY_IIISGNSSDATA_ACK struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
    GNSS_DATA_LIST []model.LocationInfoItem  
}

func (e *DOWN_EXG_MSG_APPLY_IIISGNSSDATA_ACK) GetMsgId() interface{} {
	return model.J从链路动态信息交换
}
func (e *DOWN_EXG_MSG_APPLY_IIISGNSSDATA_ACK) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *DOWN_EXG_MSG_APPLY_IIISGNSSDATA_ACK) GetDBSql() string {
	return ""
}
