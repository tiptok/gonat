package up

import (
	"github.com/tiptok/gonat/model"
)

//UP_PLATFORM_MSG_POST_QUERY_ACK 平台查岗应答
type UP_PLATFORM_MSG_POST_QUERY_ACK struct {
	model.EntityBase
    OBJECT_TYPE byte //查岗对象的类型
    OBJECT_ID string//查岗对象的ID
    INFO_ID int32// 信息ID
    INFO_CONTENT string//应答内容
    UserId string//应答用户ID
}