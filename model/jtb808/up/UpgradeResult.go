package up

import (
	"github.com/tiptok/gonat/model"
)

//UpgradeResult 终端升级结果通知 0108
type UpgradeResult struct {
	model.EntityBase
	UpgradeType byte //升级类型  byte
	Result      byte //升级结果 byte
}

func (e *UpgradeResult) GetMsgId() interface{} {
	return model.JUpgradeResult
}
func (e *UpgradeResult) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *UpgradeResult) GetDBSql() string {
	return ""
}
