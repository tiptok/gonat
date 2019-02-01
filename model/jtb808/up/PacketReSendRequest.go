package up

import (
	"github.com/tiptok/gonat/model"
)

//PacketReSendRequest 补传分包请求 (上行) 0x8003
type PacketReSendRequest struct {
	model.EntityBase
	OrginalOrderID int16 //OrginalOrderID 原始第一包序号
	PackageSNList  []int16
}

func (e *PacketReSendRequest) GetMsgId() interface{} {
	return model.JTermAuth
}
func (e *PacketReSendRequest) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *PacketReSendRequest) GetDBSql() string {
	return ""
}
