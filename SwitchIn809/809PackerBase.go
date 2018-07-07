package SwitchIn809

import (
	"bytes"

	"github.com/tiptok/gonat/model"
	"github.com/tiptok/gonat/model/jtb809/down"

	"github.com/axgle/mahonia"
	"github.com/tiptok/gotransfer/comm"
)

type JTB809PackerBase struct {
}

/*
   J1002 主链路登录应答
*/
func (p *JTB809PackerBase) J1002(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*model.UP_CONNECT_RSP)
	buf.WriteByte(inEntity.Result)
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(inEntity.Verify_Code)))
	return buf.Bytes(), nil
}

/*
   J9001 从链路连接请求
*/
func (p *JTB809PackerBase) J9001(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*model.DOWN_CONNECT_REQ)
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(inEntity.VERIFY_CODE)))
	return buf.Bytes(), nil
}

/*
   J9001 从链路连接请求
*/
func (p *JTB809PackerBase) J9005(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	//inEntity := obj.(*model.DOWN_LINKTEST_REQ)
	//buf.Write(comm.BinaryHelper.Int32ToBytes(int(inEntity.VERIFY_CODE)))
	return buf.Bytes(), nil
}

/*
   J9301 平台查岗请求
*/
func (p *JTB809PackerBase) J9301(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*down.DOWN_PLATFORM_MSG_POST_QUERY_REQ)
	enc := mahonia.NewEncoder("gbk")
	gbkdata := comm.BinaryHelper.GetASCIIString(enc.ConvertString(inEntity.INFO_CONTENT))
	buf.Write(comm.BinaryHelper.UInt16ToBytes(uint16(0x9301)))
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(len(gbkdata) + 21)))
	buf.WriteByte(inEntity.OBJECT_TYPE)
	buf.Write(comm.BinaryHelper.GetASCIIStringWL(inEntity.OBJECT_ID, 12))
	buf.Write(comm.BinaryHelper.UInt32ToBytes(uint(inEntity.INFO_ID)))
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(len(gbkdata))))
	buf.Write(gbkdata)
	return buf.Bytes(), nil
}

/*
   J9302 下发平台间报文请求
*/
func (p *JTB809PackerBase) J9302(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*down.DOWN_PLATFORM_MSG_INFO_REQ)
	enc := mahonia.NewEncoder("gbk")
	gbkdata := comm.BinaryHelper.GetASCIIString(enc.ConvertString(inEntity.INFO_CONTENT))

	buf.Write(comm.BinaryHelper.UInt16ToBytes(uint16(0x9302)))
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(len(gbkdata) + 21)))
	buf.WriteByte(inEntity.OBJECT_TYPE)
	buf.Write(comm.BinaryHelper.GetASCIIStringWL(inEntity.OBJECT_ID, 12))
	buf.Write(comm.BinaryHelper.UInt32ToBytes(uint(inEntity.INFO_ID)))
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(len(gbkdata))))
	buf.Write(gbkdata)
	return buf.Bytes(), nil
}

/*
   J9401 报警督办请求
*/
func (p *JTB809PackerBase) J9401(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*down.DOWN_WARN_MSG_URGE_TODO_REQ)
	enc := mahonia.NewEncoder("gbk")
	gbkdata := comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.Vehicle_No), 21)
	supervisor := comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.SUPERVISIOR), 16)
	buf.Write(gbkdata)
	buf.WriteByte(inEntity.Vehicle_Color)
	buf.Write(comm.BinaryHelper.UInt16ToBytes(uint16(0x9401)))
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(92)))
	buf.WriteByte(inEntity.WARN_SRC)
	buf.Write(comm.BinaryHelper.UInt16ToBytes(uint16(inEntity.WARN_TYPE)))
	buf.Write(comm.BinaryHelper.Int64ToBytes(inEntity.WARN_TIME.Unix() - 8*3600))
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(inEntity.SUPERVISION_ID)))
	buf.Write(comm.BinaryHelper.Int64ToBytes(inEntity.SUPERVISION_ENDTIME.Unix() - 8*3600))
	buf.WriteByte(inEntity.SUPERVISION_LEVEL)
	buf.Write(supervisor)
	buf.Write(comm.BinaryHelper.GetASCIIStringWL(inEntity.SUPERVISIOR_TEL, 20))
	buf.Write(comm.BinaryHelper.GetASCIIStringWL(inEntity.SUPERVISIOR_EMAIL, 32))
	return buf.Bytes(), nil
}

/*
   J9402 实时交换报警信息
*/
func (p *JTB809PackerBase) J9402(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*down.DOWN_WARN_MSG_INFORM_TIPS)
	enc := mahonia.NewEncoder("gbk")
	gbkdata := comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.Vehicle_No), 21)
	warnContent := comm.BinaryHelper.GetASCIIString(enc.ConvertString(inEntity.WARN_CONTENT))
	buf.Write(gbkdata)
	buf.WriteByte(inEntity.Vehicle_Color)
	buf.Write(comm.BinaryHelper.UInt16ToBytes(uint16(0x9402)))
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(15 + len(warnContent))))
	buf.WriteByte(inEntity.WARN_SRC)
	buf.Write(comm.BinaryHelper.UInt16ToBytes(uint16(inEntity.WARN_TYPE)))
	buf.Write(comm.BinaryHelper.Int64ToBytes(int64(inEntity.WARN_TIME.Unix() - 8*3600)))
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(len(warnContent))))
	buf.Write(warnContent)
	return buf.Bytes(), nil
}

/*
   J9403 实时交换报警信息
*/
func (p *JTB809PackerBase) J9403(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*down.DOWN_WARN_MSG_EXG_INFORM)
	enc := mahonia.NewEncoder("gbk")
	gbkdata := comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.Vehicle_No), 21)
	warnContent := comm.BinaryHelper.GetASCIIString(enc.ConvertString(inEntity.WARN_CONTENT))
	buf.Write(gbkdata)
	buf.WriteByte(inEntity.Vehicle_Color)
	buf.Write(comm.BinaryHelper.UInt16ToBytes(uint16(0x9403)))
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(15 + len(warnContent))))
	buf.WriteByte(inEntity.WARN_SRC)
	buf.Write(comm.BinaryHelper.UInt16ToBytes(uint16(inEntity.WARN_TYPE)))
	buf.Write(comm.BinaryHelper.Int64ToBytes(int64(inEntity.WARN_TIME.Unix() - 8*3600)))
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(len(warnContent))))
	buf.Write(warnContent)
	return buf.Bytes(), nil
}

// func getSubCmdCode(sSubCmdCode string)int16,error {
// 	return strconv.Atoi(sSubCmdCode)
// }

/*
   J9501 车辆单向监听请求
*/
func (p *JTB809PackerBase) J9501(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*down.DOWN_CTRL_MSG_MONITOR_VEHICLE_REQ)
	enc := mahonia.NewEncoder("gbk")
	gbkdata := comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.Vehicle_No), 21)

	buf.Write(gbkdata)
	buf.WriteByte(inEntity.Vehicle_Color)
	buf.Write(comm.BinaryHelper.UInt16ToBytes(uint16(0x9501)))
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(20)))
	buf.Write(comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.MONITOR_TEL), 20))
	return buf.Bytes(), nil
}

/*
   J9502 车辆拍照请求
*/
func (p *JTB809PackerBase) J9502(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*down.DOWN_CTRL_MSG_TAKE_PHOTO_REQ)
	enc := mahonia.NewEncoder("gbk")
	gbkdata := comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.Vehicle_No), 21)

	buf.Write(gbkdata)
	buf.WriteByte(inEntity.Vehicle_Color)
	buf.Write(comm.BinaryHelper.UInt16ToBytes(0x9502))
	buf.Write(comm.BinaryHelper.Int32ToBytes(0x02))
	buf.WriteByte(inEntity.LENS_ID)
	buf.WriteByte(inEntity.SIZE)
	return buf.Bytes(), nil
}

/*
   J9503 下发平台间报文请求
*/
func (p *JTB809PackerBase) J9503(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*down.DOWN_CTRL_MSG_TEXT_INFO)
	enc := mahonia.NewEncoder("gbk")
	plateNum := comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.Vehicle_No), 21)
	content := comm.BinaryHelper.GetASCIIString(enc.ConvertString(inEntity.MSG_CONTENT))

	buf.Write(plateNum)
	buf.WriteByte(inEntity.Vehicle_Color)
	buf.Write(comm.BinaryHelper.UInt16ToBytes(0x9503))
	buf.Write(comm.BinaryHelper.Int32ToBytes(int(len(content) + 9)))
	buf.Write(comm.BinaryHelper.Int32ToBytes(inEntity.MSG_SEQUENCE))
	buf.WriteByte(inEntity.MSG_PRIORITY)
	buf.Write(comm.BinaryHelper.Int32ToBytes(len(content)))
	buf.Write(content)
	return buf.Bytes(), nil
}

/*
   J9504 上报车辆行驶记录请求
*/
func (p *JTB809PackerBase) J9504(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*down.DOWN_CTRL_MSG_TAKE_TRAVEL_REQ)
	enc := mahonia.NewEncoder("gbk")
	plateNum := comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.Vehicle_No), 21)

	buf.Write(plateNum)
	buf.WriteByte(inEntity.Vehicle_Color)
	buf.Write(comm.BinaryHelper.UInt16ToBytes(0x9504))
	buf.Write(comm.BinaryHelper.Int32ToBytes(1))
	buf.WriteByte(inEntity.COMMAND_TYPE)
	return buf.Bytes(), nil
}

/*
   J9505 车辆应急接入监管平台请求
*/
func (p *JTB809PackerBase) J9505(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*down.DOWN_CTRL_MSG_EMERGENCY_MONITORYING_REQ)
	enc := mahonia.NewEncoder("gbk")
	plateNum := comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.Vehicle_No), 21)

	buf.Write(plateNum)
	buf.WriteByte(inEntity.Vehicle_Color)
	buf.Write(comm.BinaryHelper.UInt16ToBytes(0x9505))
	buf.Write(comm.BinaryHelper.Int32ToBytes(145))
	buf.Write(comm.BinaryHelper.Int32ToBytes(145))
	buf.Write(comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.AUTHENTICATION_CODE), 10))
	buf.Write(comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.ACCESS_POINT_NAME), 20))
	buf.Write(comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.USERNAME), 49))
	buf.Write(comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.PASSWORD), 22))
	buf.Write(comm.BinaryHelper.GetASCIIStringWL(enc.ConvertString(inEntity.SERVER_IP), 32))
	buf.Write(comm.BinaryHelper.Int16ToBytes(inEntity.TCP_PORT))
	buf.Write(comm.BinaryHelper.Int16ToBytes(inEntity.UDP_PORT))
	buf.Write(comm.BinaryHelper.Int64ToBytes(inEntity.END_TIME.Unix())) //TODO:UTC时间
	return buf.Bytes(), nil
}
