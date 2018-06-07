package core

import (
	"reflect"

	"encoding/json"

	"fmt"

	"github.com/tiptok/gonat/global"
	"github.com/tiptok/gonat/model"
	"github.com/tiptok/gotransfer/comm"
)

type bizBase struct {
	CmdType *comm.DataContext
}

//DecJSONCmd 以json格式解码下发指令
func (b bizBase) DecJSONCmd(sCmd string, cmdData []byte) (model.IEntity, error) {
	if val, ok := b.CmdType.GetOk(sCmd); ok {
		instance := reflect.New(val.(reflect.Type)).Interface()
		err := json.Unmarshal(cmdData, &instance)
		if err != nil {
			return nil, err
		}
		return instance.(model.IEntity), nil
	}
	return nil, fmt.Errorf("DecJsonCmd NOT Exists %v", sCmd)
}

//OnRegistCmdEntity  注册cmd实体
func (b bizBase) OnRegistCmdEntity() {

}
func (b bizBase) RegistCmd(sCmd string, t reflect.Type) {
	b.CmdType.Set(sCmd, t)
}

//分发接口
type IBiz interface {
	SendEntity(obj interface{}) error
	ReceiveEntity() (interface{}, error)
	Start() error
}

/*BizDBTdf 数据库分发*/
type BizDBTdf struct {
	bizBase
	TaskRec  *global.TaskManager
	TaskSend *global.TaskManager
}

//SendEntity 分发数据
func (b *BizDBTdf) SendEntity(obj interface{}) error {
	return nil
}
func (b *BizDBTdf) onSendEntity(obj interface{}) {

}

//Start 启动
func (b *BizDBTdf) Start() error {
	b.bizBase.OnRegistCmdEntity() //注册json 反序列化列表
	return nil
}

//ReceiveEntity 接收数据
func (b *BizDBTdf) ReceiveEntity() (interface{}, error) {
	return nil, nil
}
func (b *BizDBTdf) onReceiveEntity(obj interface{}) {

}

//NewBizDBTdf 新建一个BizDBTdf
func NewBizDBTdf() *BizDBTdf {
	biz := &BizDBTdf{}
	biz.bizBase = bizBase{}
	biz.TaskRec = global.NewTaskManager("BizDBTdf.Rec.Task", 1000, 1, biz.onReceiveEntity)
	biz.TaskSend = global.NewTaskManager("BizDBTdf.Send.Task", 1000, 1, biz.onReceiveEntity)
	return biz
}
