package up


import (
	"github.com/tiptok/gonat/model"
)
//UP_WARN_MSG_ADPT_TODO_INFO 主动上报报警处理结果信息消息 1403
type UP_WARN_MSG_ADPT_TODO_INFO struct{
    model.EntityBase
    Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
    INFO_ID   uint32//报警信息ID
    RESULT byte
}