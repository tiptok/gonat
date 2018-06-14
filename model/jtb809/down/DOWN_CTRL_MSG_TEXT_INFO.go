package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_CTRL_MSG_TEXT_INFO 下发车辆报文请求
type DOWN_CTRL_MSG_TEXT_INFO struct {
	model.EntityBase
	MSG_SEQUENCE int    //消息ID序号
	MSG_PRIORITY byte   // 报文优先级
	MSG_LENGTH   int    //报文信息长度
	MSG_CONTENT  string //报文信息内容
}
