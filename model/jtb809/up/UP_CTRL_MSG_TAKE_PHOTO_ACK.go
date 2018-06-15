package up

import (
	"github.com/tiptok/gonat/model"
)
//UP_CTRL_MSG_TAKE_PHOTO_ACK 车辆拍照应答 1502
type UP_CTRL_MSG_TAKE_PHOTO_ACK struct{
    model.EntityBase
    Vehicle_No    string //车牌
	Vehicle_Color byte   //车牌颜色
    PHOTO_RSP_FLAG byte //拍照应答标识
    GNSS_DATA model.LocationInfoItem//拍照位置地点
    LENS_ID byte //镜头ID
    SIZE_TYPE byte//图片大小
    TYPE byte//图像格式
    PHOTO []byte//图片内容
}