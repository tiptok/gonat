package up

import (
	"github.com/tiptok/gonat/model"
)

//TermAuth 终端鉴权 0102
type TermAuth struct {
	model.EntityBase
	AuthCode string
}

func (e *TermAuth) GetMsgId() interface{} {
	return model.JTermAuth
}
func (e *TermAuth) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *TermAuth) GetDBSql() string {
	return ""
}
