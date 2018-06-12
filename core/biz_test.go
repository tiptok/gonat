package core

import "testing"
import "log"

import "fmt"

func TestBizDBTdf(t *testing.T) {
	biz := NewBizDBTdf()
	//ctx, cancle := context.WithCancel(context.Background())
	//biz.Start(ctx)
	biz.OnRegistCmdEntity()
	//biz.ReceiveEntity()
	//time.Sleep(time.Second * 10)
	//cancle()

	checkCmd(biz, 0x9301, `{
	"MsgId":"37632",
	"AccessCode": "12345678",
	"SubMsgId": "37633",
	"OBJECT_TYPE": 2,
	"OBJECT_ID": "12345678",
	"INFO_ID": 523610,
	"INFO_CONTENT": "1+5="
}`)
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
