package up

import (
	"github.com/tiptok/gonat/model"
)

//UP_WARN_MSG_URGE_TODO_ACK 报警督办应答 1401
type UP_WARN_MSG_URGE_TODO_ACK struct {
	model.EntityBase
	Vehicle_No     string //车牌
	Vehicle_Color  byte   //车牌颜色
	SUPERVISION_ID uint32 //报警督办ID
	RESULT         byte
}
