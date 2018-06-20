package down

import (
	"time"

	"github.com/tiptok/gonat/model"
)

//DOWN_WARN_MSG_URGE_TODO_REQ 报警督办请求 0x9401
type DOWN_WARN_MSG_URGE_TODO_REQ struct {
	model.EntityBase
	Vehicle_No          string    //车牌
	Vehicle_Color       byte      //车牌颜色
	WARN_SRC            byte      //报警信息来源
	WARN_TIME           time.Time //报警时间 UTC时间格式
	WARN_TYPE           int64     //报警类型
	SUPERVISION_ID      uint32    //报警督办ＩＤ
	SUPERVISION_ENDTIME time.Time //督办截止时间，UTC时间格式
	SUPERVISION_LEVEL   byte      //督办级别
	SUPERVISIOR         string    //督办人
	SUPERVISIOR_TEL     string    //督办联系电话
	WARN_CONTENT        string    //报警描述
	SUPERVISIOR_EMAIL   string    //督办联系电子邮件
}
