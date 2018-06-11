package core

import "testing"
import "log"

func TestBizDBTdf(t *testing.T) {
	biz := NewBizDBTdf()
	//biz.Start()
	biz.OnRegistCmdEntity()
	e, err := biz.DecCmd("json", "37633", []byte(`{   "OBJECT_TYPE":2,   "OBJECT_ID":"12345678",   "INFO_ID":523610,   "INFO_CONTENT":"1+5="  }`))
	if err != nil {
		log.Println(err)
	}
	if e != nil {

	}
}
