package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_CTRL_MSG_TAKE_PHOTO_REQ 车辆拍照请求 0x9502
type DOWN_CTRL_MSG_TAKE_PHOTO_REQ struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
	LENS_ID       byte   //镜头ID
	SIZE          byte   //SIZE
}

func (e *DOWN_CTRL_MSG_TAKE_PHOTO_REQ) GetMsgId() interface{} {
	return model.J从链路车辆监管
}
func (e *DOWN_CTRL_MSG_TAKE_PHOTO_REQ) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *DOWN_CTRL_MSG_TAKE_PHOTO_REQ) GetDBSql() string {
	return ""
}
