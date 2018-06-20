package up


import (
	"github.com/tiptok/gonat/model"
	"time"
)
//UP_WARN_MSG_ADPT_INFO 上报报警信息 1402
type UP_WARN_MSG_ADPT_INFO struct{
    model.EntityBase
    Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
    WARN_SRC   byte//报警信息ID
    WARN_TYPE  int64//报警类型
    WARN_TIME  time.Time//报警时间 UTC时间格式
    INFO_ID    uint
    INFO_CONTENT string//报警描述
}