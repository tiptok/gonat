package down

import (
	"github.com/tiptok/gonat/model"
)

//DOWN_WARN_MSG 从链路平台间信息交换
type DOWN_WARN_MSG struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
}
