package core

import (
	"fmt"
	"log"
	"testing"

	"encoding/json"

	"github.com/axgle/mahonia"
	"github.com/tiptok/gonat/model/jtb809/down"
)

func TestBizDBTdf(t *testing.T) {
	biz := NewBizDBTdf()
	//ctx, cancle := context.WithCancel(context.Background())
	//biz.Start(ctx)
	biz.OnRegistCmdEntity()
	//biz.ReceiveEntity()
	//time.Sleep(time.Second * 10)
	//cancle()

	checkCmd(biz, 0x9302, `{
	"MsgId":"37632",
	"AccessCode": "12345678",
	"SubMsgId": "37633",
	"OBJECT_TYPE": 2,
	"OBJECT_ID": "12345678",
	"INFO_ID": 523610,
	"INFO_CONTENT": "1+5="
}`)
	checkCmd(biz, 0x9301, `{"MsgId":37632,"MsgSN":0,"SubMsgId":37633,"AccessCode":"12345678","OBJECT_TYPE":2,"OBJECT_ID":"12345678","INFO_ID":0,"INFO_CONTENT":"1+3"}`)
	e9301 := down.DOWN_PLATFORM_MSG_POST_QUERY_REQ{}
	e9301.OBJECT_ID = "12345678"
	checkCmdJson(e9301)
}

func checkCmd(biz *BizDBTdf, cmdCode int, cmdData string) {
	e, err := biz.DecCmd("json", fmt.Sprintf("%d", cmdCode), []byte(cmdData))
	if err != nil {
		log.Println(err)
	}
	if e != nil {
		log.Printf("%x Entity:%v", cmdCode, e)
	}
}

func checkCmdJson(obj interface{}) {
	data, err := json.Marshal(obj)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Deccmd :%s", string(data))
}

func TestEncodeGbk(t *testing.T) {
	//"你好，世界！"的GBK编码
	testBytes := []byte{0xC4, 0xE3, 0xBA, 0xC3, 0xA3, 0xAC, 0xCA, 0xC0, 0xBD, 0xE7, 0xA3, 0xA1}
	var testStr string
	utfStr := "你好，世界！"
	var dec mahonia.Decoder
	var enc mahonia.Encoder

	testStr = string(testBytes)

	dec = mahonia.NewDecoder("gbk")
	fmt.Println("GBK", []byte(utfStr))
	ret := dec.ConvertString(testStr)
	fmt.Println("GBK to UTF-8: ", ret, " bytes:", testBytes)

	enc = mahonia.NewEncoder("gbk")
	ret = enc.ConvertString(utfStr)
	fmt.Println("UTF-8 to GBK: ", ret, " bytes: ", []byte(ret))
}
