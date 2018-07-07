package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_EXG_MSG_APPLY_FOR_MONITOR_STARTUP_ACK 申请交换指定车辆定位信息应答 9207
type DOWN_EXG_MSG_APPLY_FOR_MONITOR_STARTUP_ACK struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
	/*  
    /// 取消交换指定车辆定位信息结果 
    /// 0:申请成功
    /// 1：上级没有该车辆信息
    /// 2:申请时间错误
    /// 3:其他
    */
    Result  byte   
}

func (e *DOWN_EXG_MSG_APPLY_FOR_MONITOR_STARTUP_ACK) GetMsgId() interface{} {
	return model.J从链路动态信息交换
}
func (e *DOWN_EXG_MSG_APPLY_FOR_MONITOR_STARTUP_ACK) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *DOWN_EXG_MSG_APPLY_FOR_MONITOR_STARTUP_ACK) GetDBSql() string {
	return ""
}
