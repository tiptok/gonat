package SwitchIn809

import (
	"bytes"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"runtime/debug"
	"strconv"

	"github.com/tiptok/gonat/global"
	"github.com/tiptok/gonat/model"

	"github.com/tiptok/gotransfer/comm"
	"github.com/tiptok/gotransfer/conn"
)

//var _JTB808ParseBase *JTB808ParseBase = &JTB808ParseBase{}
type protocol809 struct {
}

func (p protocol809) PacketMsg(obj interface{}) (data []byte, err error) {
	defer func() {
		if p := recover(); p != nil {
			debug.PrintStack()
			log.Printf("protocol809.PacketMsg panic recover! p: %v", p)
		}
	}()
	packdata, err := p.Packet(obj)
	if err != nil {
		return nil, err
	}

	if def, ok := obj.(model.IEntity); ok {
		entity := def.GetEntityBase()
		if packdata != nil && len(packdata) > 0 {
			global.Debug("MsgId:%X MsgBodyData:%s", def.GetMsgId().(int), comm.BinaryHelper.ToBCDString(packdata, 0, int32(len(packdata))))
		}
		buf := bytes.NewBuffer(nil)
		buf.Write(comm.BinaryHelper.Int32ToBytes(len(packdata) + 26))          //总长度
		buf.Write(comm.BinaryHelper.Int32ToBytes(entity.MsgSN))                //流水号
		buf.Write(comm.BinaryHelper.Int16ToBytes(int16(def.GetMsgId().(int)))) //消息Id
		iAccesscode, _ := strconv.Atoi(entity.AccessCode)
		buf.Write(comm.BinaryHelper.Int32ToBytes(iAccesscode)) //接入码
		buf.Write([]byte{0x00, 0x01, 0x00})                    //版本号 0.1.0
		if global.Param.IsEncrypt {                            //是否加密
			buf.WriteByte(0x01)
		} else {
			buf.WriteByte(0x00)
		}
		buf.Write(comm.BinaryHelper.Int32ToBytes(global.Param.Key)) //密钥
		buf.Write(packdata)
		crc := comm.BinaryHelper.CRC16Check(buf.Bytes()) //计算crc
		buf.Write(comm.BinaryHelper.Int16ToBytes(crc))
		return Byte809Enscape(buf.Bytes(), 0, buf.Len()), nil
	}
	return packdata, err
}

/*
	打包数据体
	obj 数据体
*/
func (p protocol809) Packet(obj interface{}) (packdata []byte, err error) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("protocol809.Packet panic recover! p: %v", p)
		}
	}()
	if def, ok := obj.(model.IEntity); ok {
		entity := def.GetEntityBase()
		var sMethodName string
		var cmdCode interface{}
		if entity.SubMsgId == nil {
			cmdCode = entity.MsgId
		} else {
			cmdCode = entity.SubMsgId
		}
		icmdCode, err := strconv.Atoi(fmt.Sprintf("%v", cmdCode))
		if err != nil {
			return nil, err
			//return nil, fmt.Errorf("Packet:cmdCode type is not int %v %v", cmdCode, reflect.TypeOf(cmdCode))
		}
		sMethodName = fmt.Sprintf("J%X", icmdCode)
		global.Debug("InvokeFunc:%s", sMethodName)
		value, err := comm.ParseHelper.InvokeFunc(&JTB809PackerBase{}, sMethodName, obj)
		if err == nil {
			// if value == nil || len(value) == 0 {
			// 	return nil,errors.New("protocol809.Packet:not pakcage return from 809PackerBase.")
			// }
			packdata = (value[0].Interface()).([]byte)
		} else {
			return nil, err
		}
	} else {
		err = errors.New("非809实体,发送异常")
	}
	return packdata, err
}

/*
	分包处理
	packdata 解析出一个完整包
	leftdata 解析剩余报文的数据
	err 	 分包错误
*/
func (p protocol809) ParseMsg(data []byte, c *conn.Connector) (packdata [][]byte, leftdata []byte, err error) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("protocol809.ParseMsg panic recover! p: %v", p)
		}
	}()
	packdata, leftdata, err = comm.ParseHelper.ParsePartWithPool(data, 0x5b, 0x5d, c.Pool)
	if err == nil {
		//global.Debug(global.F(global.TCP, global.SVR809, "接收到数据包 : %v"), hex.EncodeToString(data))
		// for i, v := range packdata {
		// 	global.Debug(global.F(global.TCP, global.SVR809, "序号%d : %v"), i, hex.EncodeToString(v))
		// }

		//global.Debug(global.F(global.TCP, global.SVR809, "ParseMsg Package size:(%X)  LeftData:%v(%d)"), len(packdata), comm.BinaryHelper.ToBCDString(leftdata, int32(0), int32(len(leftdata))), len(leftdata))

	}
	return packdata, leftdata, err
}

/*
	解析数据
	obj 解析出对应得数据结构
	err 解析出错
*/
func (p protocol809) Parse(packdata []byte) (obj interface{}, err error) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("protocol809.Parse panic recover! p: %v", p)
		}
	}()
	//转义
	data, err := Byte809Descape(packdata, 0, len(packdata))

	global.Debug(global.F(global.TCP, global.SVR809, "收到数据: %v"), hex.EncodeToString(packdata))
	//CRC Check
	tmpCrc := comm.BinaryHelper.ToInt16(data, int32(len(data)-2))
	checkCrc := comm.BinaryHelper.CRC16Check(data[:len(data)-2])
	if uint16(checkCrc) != uint16(tmpCrc) {
		err = errors.New(fmt.Sprintf("CRC CHECK Error->%x != %x %v", uint16(tmpCrc), uint16(checkCrc), hex.EncodeToString(packdata)))
		return nil, err
	}
	h := model.EntityBase{}
	//数据头
	msgBodyLenght := comm.BinaryHelper.ToInt32(data, 0) - 26
	h.MsgSN = int(comm.BinaryHelper.ToInt32(data, 4))
	h.MsgId = comm.BinaryHelper.ToUInt16(data, 8)
	h.AccessCode = fmt.Sprintf("%d", comm.BinaryHelper.ToInt32(data, 10))

	isEncrypt := data[17] == 0
	if isEncrypt {
		//解密
	}
	msgBody := data[22 : msgBodyLenght+22]

	sMethodName := fmt.Sprintf("J%s", comm.BinaryHelper.ToBCDString(data, 8, 2))
	//InvokeFunc
	//log.Println(msgBodyLenght, sMethodName, comm.BinaryHelper.ToBCDString(msgBody, 0, int32(len(msgBody))))
	value, err := comm.ParseHelper.InvokeFunc(&JTB809ParseBase{}, sMethodName, msgBody, h)
	if err != nil {
		return nil, err
	}
	if value != nil && len(value) > 0 {
		obj = value[0].Interface()
	}

	return obj, err
}
