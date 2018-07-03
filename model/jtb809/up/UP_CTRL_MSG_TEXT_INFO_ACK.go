package up

import (
	"github.com/tiptok/gonat/model"
)

//UP_CTRL_MSG_TEXT_INFO_ACK 下发车辆报文应答 1503
type UP_CTRL_MSG_TEXT_INFO_ACK struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
	MSG_ID        uint32 //消息ID
	RESULT        byte   //下发车辆报文应答结果
}
