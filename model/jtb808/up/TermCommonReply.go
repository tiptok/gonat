package up

import (
	"github.com/tiptok/gonat/model"
)

//TermHeartbeat 终端心跳 8100
type TermCommonReply struct {
	model.EntityBase
	RspMsgSN  int
	RspMsgId  int
	RspResult byte
}

func (e *TermCommonReply) GetMsgId() interface{} {
	return model.JTermCommonReply
}
func (e *TermCommonReply) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *TermCommonReply) GetDBSql() string {
	return ""
}
