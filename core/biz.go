package core

import (
	"reflect"

	"encoding/json"

	"fmt"

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

//BizDBTdf 数据库分发
type BizDBTdf struct {
	bizBase
}
