package SwitchIn809

import (
	"fmt"
	"strings"
	"time"

	"github.com/axgle/mahonia"
	"github.com/tiptok/gonat/model"
	"github.com/tiptok/gonat/model/jtb809/up"
	"github.com/tiptok/gotransfer/comm"
)

type JTB809ParseBase struct {
}

// func(p *JTB809ParseBase)Parse(h MsgHead,msgBody []byte)(interface{},error){

// }

//J1001 0x1001 主链路登录请求
func (p *JTB809ParseBase) J1001(msgBody []byte, h model.EntityBase) interface{} {
	outEntity := &model.UP_CONNECT_REQ{}
	outEntity.SetEntity(h)
	outEntity.USERID = uint32(comm.BinaryHelper.ToInt32(msgBody, 0)) //应答ID
	outEntity.PASSWORD = comm.BinaryHelper.ToASCIIString(msgBody, 4, 8)
	outEntity.DOWN_LINK_IP = strings.Trim(comm.BinaryHelper.ToASCIIString(msgBody, 12, 32), string([]byte{0x00})) //结果
	outEntity.DOWN_LINK_PORT = uint32(comm.BinaryHelper.ToInt16(msgBody, 44))                                     //应答ID

	/*是否需要应答*/
	// outEntity.IsNeedRsp = false
	// outEntity.DownRspEntity = nil
	return outEntity
}

//J1002 主链路登录应答 0x1002
func (p *JTB809ParseBase) J1002(msgBody []byte, h model.EntityBase) interface{} {
	outEntity := &model.UP_CONNECT_RSP{}
	outEntity.SetEntity(h)
	outEntity.Result = msgBody[0]
	outEntity.Verify_Code = comm.BinaryHelper.ToInt32(msgBody, 1)
	/*是否需要应答*/
	// outEntity.IsNeedRsp = false
	// outEntity.DownRspEntity = nil
	return outEntity
}

//J9002 从链路登录应答
func (p *JTB809ParseBase) J9002(msgBody []byte, h model.EntityBase) interface{} {
	outEntity := &model.DOWN_CONNECT_RSP{}
	outEntity.SetEntity(h)
	outEntity.Result = msgBody[0]
	/*是否需要应答*/
	// outEntity.IsNeedRsp = false
	// outEntity.DownRspEntity = nil
	return outEntity
}

//J1005 0x1005  主链路保持连接请求
func (p *JTB809ParseBase) J1005(msgBody []byte, h model.EntityBase) interface{} {
	outEntity := &model.UP_LINKTEST_REQ{}
	outEntity.SetEntity(h)
	// outEntity.Result = msgBody[0]
	// outEntity.Verify_Code = comm.BinaryHelper.ToInt32(msgBody, 1)
	/*是否需要应答*/
	// outEntity.IsNeedRsp = false
	// outEntity.DownRspEntity = nil
	return outEntity
}

//J9006 从链路登录应答
func (p *JTB809ParseBase) J9006(msgBody []byte, h model.EntityBase) interface{} {
	outEntity := &model.DOWN_LINKTEST_RSP{}
	outEntity.SetEntity(h)
	return outEntity
}

//J1200 主链路动态信息交换
func (p *JTB809ParseBase) J1200(msgBody []byte, h model.EntityBase) interface{} {
	outEntity := model.UP_EXG_MSG{}
	outEntity.SetEntity(h)
	enc := mahonia.NewDecoder("gbk")
	outEntity.Vehicle_No = strings.Trim(comm.BinaryHelper.ToASCIIString(msgBody, 0, 21), string([]byte{0x00})) // strings.Trim(comm.BinaryHelper.ToASCIIString(msgBody, 0, 21), " ")
	outEntity.Vehicle_No = enc.ConvertString(outEntity.Vehicle_No)
	outEntity.Vehicle_Color = msgBody[21]
	outEntity.SubMsgId = comm.BinaryHelper.ToUInt16(msgBody, 22)
	switch outEntity.SubMsgId.(uint16) {
	case 0x1201:
		return J1201(msgBody, outEntity)
	case 0x1202:
		return J1202(msgBody, outEntity)
	case 0x1203:
		return J1203(msgBody, outEntity)
	default:
		panic(fmt.Sprintf("未找到对应方法:%v", outEntity.SubMsgId))
	}
	return nil
}
func J1201(msgBody []byte, h model.UP_EXG_MSG) interface{} {
	outEntity := &up.UP_EXG_MSG_REGISTER{
		UP_EXG_MSG: h,
	}

	return outEntity
}
func J1202(msgBody []byte, h model.UP_EXG_MSG) interface{} {
	outEntity := &model.UP_EXG_MSG_REAL_LOCATION{
		UP_EXG_MSG: h,
	}
	outEntity.GNSS_DATA = GetGetLoactionInfo(msgBody, 28)
	return outEntity
}
func J1203(msgBody []byte, h model.UP_EXG_MSG) interface{} {
	outEntity := &model.UP_EXG_MSG_HISTORY_LOCATION{
		UP_EXG_MSG: h,
	}
	//iIndex := 29
	posSize := int(msgBody[28])
	outEntity.GNSS_DATA_LIST = make([]model.LocationInfoItem, posSize)
	for i := 0; i < posSize; i++ {
		outEntity.GNSS_DATA_LIST[i] = GetGetLoactionInfo(msgBody, int32(29+i*36))
	}
	return outEntity
}

func GetGetLoactionInfo(msgBody []byte, iIndex int32) (location model.LocationInfoItem) {
	location.ENCRYPT = msgBody[iIndex]
	iIndex += 1
	iDay := int(msgBody[iIndex])
	iMonth := int(msgBody[iIndex+1])
	iYear := int(comm.BinaryHelper.ToInt16(msgBody, int32(iIndex+2)))
	iHour := int(msgBody[iIndex+4])
	iMin := int(msgBody[iIndex+5])
	iSec := int(msgBody[iIndex+6])
	var err error
	//location.GPSTIME, err = time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%v-%v-%v %v:%v:%v", iYear, iMonth, iDay, iHour, iMin, iSec))
	location.GPSTIME = time.Date(iYear, time.Month(iMonth), iDay, iHour, iMin, iSec, 0, time.Local)
	if err != nil {
		panic(err)
	}
	location.LON = int(comm.BinaryHelper.ToInt32(msgBody, iIndex+7))
	location.LAT = int(comm.BinaryHelper.ToInt32(msgBody, iIndex+11))
	location.VEC1 = comm.BinaryHelper.ToInt16(msgBody, iIndex+15)
	location.VEC2 = comm.BinaryHelper.ToInt16(msgBody, iIndex+17)
	location.VEC3 = int(comm.BinaryHelper.ToInt32(msgBody, iIndex+19))
	location.DIRECTION = comm.BinaryHelper.ToInt16(msgBody, iIndex+23)
	location.ALTITUDE = comm.BinaryHelper.ToInt16(msgBody, iIndex+25)
	location.STATE = int(comm.BinaryHelper.ToInt32(msgBody, iIndex+27))
	location.ALARM = int(comm.BinaryHelper.ToInt32(msgBody, iIndex+31))
	return location
}

//J1200 主链路动态信息交换
func (p *JTB809ParseBase) J1300(msgBody []byte, h model.EntityBase) interface{} {
	outEntity := model.EntityBase{}
	outEntity.SetEntity(h)
	outEntity.SubMsgId = comm.BinaryHelper.ToUInt16(msgBody, 0)
	switch outEntity.SubMsgId.(uint16) {
	case 0x1301:
		return J1301(msgBody, outEntity)
	case 0x1302:
		return J1302(msgBody, outEntity)
	default:
		panic(fmt.Sprintf("未找到对应方法:%v", outEntity.SubMsgId))
	}
	return nil
}

func J1301(msgBody []byte, h model.EntityBase) interface{} {
	outEntity := &up.UP_PLATFORM_MSG_POST_QUERY_ACK{
		EntityBase: h,
	}
	//length := comm.BinaryHelper.ToInt32(msgBody, 2)
	outEntity.OBJECT_TYPE = msgBody[6] //报警信息来源
	outEntity.OBJECT_ID = comm.BinaryHelper.ToASCIIString(msgBody, 7, 12)
	outEntity.INFO_ID = comm.BinaryHelper.ToInt32(msgBody, 19)
	dec := mahonia.NewDecoder("gbk")
	lenContent := comm.BinaryHelper.ToInt32(msgBody, 23)
	outEntity.INFO_CONTENT = dec.ConvertString(string(msgBody[23 : 23+lenContent]))
	return outEntity
}
func J1302(msgBody []byte, h model.EntityBase) interface{} {
	outEntity := &up.UP_PLATFORM_MSG_INFO_ACK{
		EntityBase: h,
	}
	//4
	outEntity.INFO_ID = comm.BinaryHelper.ToInt32(msgBody, 6)
	return outEntity
}

//J1500 主链路车辆监管
func (p *JTB809ParseBase) J1500(msgBody []byte, h model.EntityBase) interface{} {
	outEntity := up.UP_CTRL_MSG{}
	outEntity.SetEntity(h)
	outEntity.SubMsgId = comm.BinaryHelper.ToUInt16(msgBody, 0)
	dec := mahonia.NewDecoder("gbk")
	outEntity.Vehicle_No = strings.Trim(comm.BinaryHelper.ToASCIIString(msgBody, 0, 21), string([]byte{0x00})) // strings.Trim(comm.BinaryHelper.ToASCIIString(msgBody, 0, 21), " ")
	outEntity.Vehicle_No = dec.ConvertString(outEntity.Vehicle_No)
	outEntity.Vehicle_Color = msgBody[21]
	switch outEntity.SubMsgId.(uint16) {
	case 0x1501:
		return J1501(msgBody, outEntity)
	// case 0x1502:
	// 	return J1502(msgBody, outEntity)
	default:
		panic(fmt.Sprintf("未找到对应方法:%v", outEntity.SubMsgId))
	}
	return nil
}

//J1501 车辆单向监听应答
func J1501(msgBody []byte, h up.UP_CTRL_MSG) interface{} {
	outEntity := &up.UP_CTRL_MSG_MONITOR_VEHICLE_ACK{
		EntityBase:    h.EntityBase,
		Vehicle_No:    h.Vehicle_No,
		Vehicle_Color: h.Vehicle_Color,
	}
	//lenght 4
	outEntity.RESULT = msgBody[28]
	return outEntity
}

//J1502 车辆拍照应答
func J1502(msgBody []byte, h up.UP_CTRL_MSG) interface{} {
	outEntity := &up.UP_CTRL_MSG_TAKE_PHOTO_ACK{
		EntityBase:    h.EntityBase,
		Vehicle_No:    h.Vehicle_No,
		Vehicle_Color: h.Vehicle_Color,
	}
	//lenght 4
	outEntity.PHOTO_RSP_FLAG = msgBody[28]
	outEntity.GNSS_DATA = GetGetLoactionInfo(msgBody, 29)
	outEntity.LENS_ID = msgBody[65]
	lenPhoto := comm.BinaryHelper.ToInt32(msgBody, 66)
	outEntity.TYPE = msgBody[70]
	outEntity.PHOTO = msgBody[71 : 71+lenPhoto-1]
	return outEntity
}

//J1503 下发车辆报文应答
func J1503(msgBody []byte, h up.UP_CTRL_MSG) interface{} {
	outEntity := &up.UP_CTRL_MSG_TEXT_INFO_ACK{
		EntityBase:    h.EntityBase,
		Vehicle_No:    h.Vehicle_No,
		Vehicle_Color: h.Vehicle_Color,
	}
	//lenght 4
	outEntity.MSG_ID = comm.BinaryHelper.ToUInt32(msgBody, 28)
	outEntity.RESULT = msgBody[32]
	return outEntity
}

//J1504 上报车辆行驶记录应答
func J1504(msgBody []byte, h up.UP_CTRL_MSG) interface{} {
	outEntity := &up.UP_CTRL_MSG_TAKE_TRAVEL_ACK{
		EntityBase:    h.EntityBase,
		Vehicle_No:    h.Vehicle_No,
		Vehicle_Color: h.Vehicle_Color,
	}
	//lenght 4
	lenData := comm.BinaryHelper.ToInt32(msgBody, 32)
	outEntity.TRAVELDATA_INFO = comm.BinaryHelper.ToBCDString(msgBody, 36, lenData)
	return outEntity
}

//J1505  车辆应急接入督管平台应答
func J1505(msgBody []byte, h up.UP_CTRL_MSG) interface{} {
	outEntity := &up.UP_CTRL_MSG_EMERGENCY_MONITORING_ACK{
		EntityBase:    h.EntityBase,
		Vehicle_No:    h.Vehicle_No,
		Vehicle_Color: h.Vehicle_Color,
	}
	//lenght 4
	//outEntity.MSG_ID = comm.BinaryHelper.ToUInt32(msgBody, 28)
	outEntity.RESULT = msgBody[28]
	return outEntity
}
