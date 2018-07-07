package up

import "github.com/tiptok/gonat/model"

//UP_EXG_MSG_REAL_LOCATION 补发车辆定位信息请求 1202
type UP_EXG_MSG_REAL_LOCATION struct {
	model.EntityBase
	Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
	GNSS_DATA     model.LocationInfoItem
}
