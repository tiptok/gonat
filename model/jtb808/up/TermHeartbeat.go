package up

import (
	"github.com/tiptok/gonat/model"
)

//TermHeartbeat 终端心跳 8100
type TermHeartbeat struct {
	model.EntityBase
}

func (e *TermHeartbeat) GetMsgId() interface{} {
	return model.JTermHeartbeat
}
func (e *TermHeartbeat) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *TermHeartbeat) GetDBSql() string {
	return ""
}
