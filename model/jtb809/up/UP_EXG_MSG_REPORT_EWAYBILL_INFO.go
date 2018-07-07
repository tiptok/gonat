package up

import "github.com/tiptok/gonat/model"

//UP_EXG_MSG_REPORT_DRIVER_INFO 主动上报驾驶员身份信息消息 120C
type UP_EXG_MSG_REPORT_EWAYBILL_INFO struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
	EWAYBILL_INFO string //电子运单数据内容
	//VALID_DATE    time.Time //证件有效期
}

func (e *UP_EXG_MSG_REPORT_EWAYBILL_INFO) GetMsgId() interface{} {
	return model.J主链路动态信息交换
}
func (e *UP_EXG_MSG_REPORT_EWAYBILL_INFO) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *UP_EXG_MSG_REPORT_EWAYBILL_INFO) GetDBSql() string {
	return ""
}
