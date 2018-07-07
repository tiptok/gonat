package up

import "github.com/tiptok/gonat/model"

//UP_EXG_MSG_HISTORY_LOCATION 补发车辆定位信息请求 1203
type UP_EXG_MSG_HISTORY_LOCATION struct {
	model.EntityBase
	Vehicle_No     string //车牌
	Vehicle_Color  byte   //车牌颜色
	GNSS_DATA_LIST []model.LocationInfoItem
}
