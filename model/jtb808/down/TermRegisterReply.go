package down

import (
	"github.com/tiptok/gonat/model"
)

//TermRegisterReply 终端注册应答 8100
type TermRegisterReply struct {
	model.EntityBase
	RspResult byte
}

func (e *TermRegisterReply) GetMsgId() interface{} {
	return model.JTermRegisterReply
}
func (e *TermRegisterReply) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *TermRegisterReply) GetDBSql() string {
	return ""
}
