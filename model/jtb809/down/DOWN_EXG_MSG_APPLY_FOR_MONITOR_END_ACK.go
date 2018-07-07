package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_EXG_MSG_APPLY_FOR_MONITOR_END_ACK 取消交换指定车辆定位信息应答 9208
type DOWN_EXG_MSG_APPLY_FOR_MONITOR_END_ACK struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
	/*  
    /// 取消交换指定车辆定位信息结果 
    /// 0:取消申请成功
    /// 1：之前没有对应申请信息
    /// 2:其他
    */
    Result  byte   
}

func (e *DOWN_EXG_MSG_APPLY_FOR_MONITOR_END_ACK) GetMsgId() interface{} {
	return model.J从链路动态信息交换
}
func (e *DOWN_EXG_MSG_APPLY_FOR_MONITOR_END_ACK) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *DOWN_EXG_MSG_APPLY_FOR_MONITOR_END_ACK) GetDBSql() string {
	return ""
}
