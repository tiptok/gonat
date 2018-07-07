package up

import "github.com/tiptok/gonat/model"

//UP_EXG_MSG_RETURN_STARTUP_ACK 结束车辆定位信息交换应答 1205
type UP_EXG_MSG_RETURN_STARTUP_ACK struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
	//EWAYBILL_INFO string //电子运单数据内容
	//VALID_DATE    time.Time //证件有效期
}

func (e *UP_EXG_MSG_RETURN_STARTUP_ACK) GetMsgId() interface{} {
	return model.J主链路动态信息交换
}
func (e *UP_EXG_MSG_RETURN_STARTUP_ACK) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *UP_EXG_MSG_RETURN_STARTUP_ACK) GetDBSql() string {
	return ""
}
