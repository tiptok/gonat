package up

import (
	"github.com/tiptok/gonat/model"
)

//UP_EXG_MSG_REGISTER 上传车辆注册信息 1201
type UP_EXG_MSG_REGISTER struct {
	model.UP_EXG_MSG
	PlatForm_Id     string //平台唯一编码
	Produce_Id      string //车载终端厂商唯一编码
	Term_Model_Type string //终端型号 不足8位 以‘\0’终结
	Term_Id         string //终端编号 大写字母和数字组成
	Term_SimCode    string //车载终端SIM卡电话号码 。号码不足12位，则在前补充数字0
}

func (e *UP_EXG_MSG_REGISTER) GetMsgId() interface{} {
	return model.J主链路动态信息交换
}
func (e *UP_EXG_MSG_REGISTER) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *UP_EXG_MSG_REGISTER) GetDBSql() string {
	return ""
}
