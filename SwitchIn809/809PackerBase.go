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
   J9301 从链路连接请求
*/
func (p *JTB809PackerBase) J9301(obj interface{}) (packdata []byte, err error) {
	buf := bytes.NewBuffer(nil)
	inEntity := obj.(*down.DOWN_PLATFORM_MSG_POST_QUERY_REQ)
	enc := mahonia.NewDecoder("gbk")
	gbkdata := comm.BinaryHelper.GetASCIIString(enc.ConvertString(inEntity.INFO_CONTENT))

	buf.Write(gbkdata)
	//buf.Write(comm.BinaryHelper.Int32ToBytes(int(inEntity.VERIFY_CODE)))
	return buf.Bytes(), nil
}
func getSubCmdCode(sSubCmdCode string) {

}
