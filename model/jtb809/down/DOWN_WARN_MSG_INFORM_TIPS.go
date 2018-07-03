package down

import (
	"time"

	"github.com/tiptok/gonat/model"
)

//DOWN_WARN_MSG_INFORM_TIPS 实时交换报警信息 0x9402
type DOWN_WARN_MSG_INFORM_TIPS struct {
	model.EntityBase
	Vehicle_No    string    //车牌
	Vehicle_Color byte      //车牌颜色
	WARN_SRC      byte      //报警信息来源
	WARN_TIME     time.Time //报警时间 UTC时间格式
	WARN_TYPE     int64     //报警类型
	WARN_CONTENT  string    //报警描述
}
