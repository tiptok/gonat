package up

import (
	"github.com/tiptok/gonat/model"
)

//UP_PLATFORM_MSG_INFO_ACK 下发平台间报文应答
type UP_PLATFORM_MSG_INFO_ACK struct {
	model.EntityBase
	INFO_ID int32
}
