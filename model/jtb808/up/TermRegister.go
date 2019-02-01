package up

import (
	"github.com/tiptok/gonat/model"
)

//TermRegisterReply 终端注册应答 8100
type TermRegister struct {
	model.EntityBase
	ProvinceId     int
	CityId         int
	ManufacturerId string
	TerminalModel  string
	TerminalId     string
	PlateColor     int
	PlateNum       string
}

func (e *TermRegister) GetMsgId() interface{} {
	return model.JTermRegister
}
func (e *TermRegister) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *TermRegister) GetDBSql() string {
	return ""
}
