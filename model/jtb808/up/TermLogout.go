package up

import (
	"github.com/tiptok/gonat/model"
)

//TermHeartbeat 终端注销0x0003
type TermLogout struct {
	model.EntityBase
}

func (e *TermLogout) GetMsgId() interface{} {
	return model.JTermLogout
}
func (e *TermLogout) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *TermLogout) GetDBSql() string {
	return ""
}
