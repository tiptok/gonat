package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_CTRL_MSG_TAKE_PHOTO_REQ 车辆拍照请求
type DOWN_CTRL_MSG_TAKE_PHOTO_REQ struct {
	model.EntityBase
	LENS_ID byte //镜头ID
	SIZE    byte //SIZE
}
