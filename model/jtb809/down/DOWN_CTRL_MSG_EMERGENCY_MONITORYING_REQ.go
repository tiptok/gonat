package down

import (
	"time"

	"github.com/tiptok/gonat/model"
)

//DOWN_CTRL_MSG_EMERGENCY_MONITORYING_REQ 车辆应急接入监管平台请求 0x1505
type DOWN_CTRL_MSG_EMERGENCY_MONITORYING_REQ struct {
	model.EntityBase
	Vehicle_No          string    //车牌
	Vehicle_Color       byte      //车牌颜色
	AUTHENTICATION_CODE string    //监管平台下发的鉴权码
	ACCESS_POINT_NAME   string    //拨号点名称
	USERNAME            string    //拨号用户名
	PASSWORD            string    //拨号密码
	SERVER_IP           string    // 地址 服务器IP
	TCP_PORT            int16     //服务器 TCP端口
	UDP_PORT            int16     //服务器 UDP端口
	END_TIME            time.Time //结束时间
}
