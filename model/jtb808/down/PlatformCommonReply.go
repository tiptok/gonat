package down

import (
	"github.com/tiptok/gonat/model"
)

//PlatformCommonReply 平台通用应答 8001
type PlatformCommonReply struct {
	model.EntityBase
	RspMsgSN  int
	RspMsgId  int
	RspResult byte // 应答结果，0：成功/确认；1：失败；2：消息有误；3：不支持；4：报警处理确认
}

func (e *PlatformCommonReply) GetMsgId() interface{} {
	return model.JPlatformCommonReply
}
func (e *PlatformCommonReply) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *PlatformCommonReply) GetDBSql() string {
	return ""
}
