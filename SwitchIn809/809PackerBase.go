package SwitchIn809

import (
	"bytes"

	"github.com/axgle/mahonia"
	"github.com/tiptok/gonat/model"
	"github.com/tiptok/gonat/model/jtb809/down"
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

// func getSubCmdCode(sSubCmdCode string)int16,error {
// 	return strconv.Atoi(sSubCmdCode)
// }
