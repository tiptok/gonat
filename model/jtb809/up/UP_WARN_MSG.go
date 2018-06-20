package up

import (
	"github.com/tiptok/gonat/model"
)

//UP_WARN_MSG 主链路报警信息交互 0x1400
type UP_WARN_MSG struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
}