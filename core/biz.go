package core

import (
	"context"
	"log"
	"reflect"

	"encoding/json"

	"fmt"

	"time"

	"github.com/tiptok/gonat/global"
	"github.com/tiptok/gonat/model"
	"github.com/tiptok/gonat/model/jtb809/down"
	"github.com/tiptok/gotransfer/comm"
)

type bizBase struct {
	CmdType *comm.DataContext
}

func newbizBase() bizBase {
	return bizBase{
		CmdType: &comm.DataContext{},
	}
}
func (b bizBase) DecCmd(decType, sCmd string, cmdData []byte) (model.IEntity, error) {
	switch decType {
	case "json":
		return b.DecJSONCmd(sCmd, cmdData)
	}
	return nil, nil
}

//DecJSONCmd 以json格式解码下发指令
func (b bizBase) DecJSONCmd(sCmd string, cmdData []byte) (model.IEntity, error) {
	if b.CmdType == nil {
		return nil, fmt.Errorf("bizBase.CmdType is nil %v", sCmd)
	}
	if val, ok := b.CmdType.GetOk(sCmd); ok {
		instance := reflect.New(val.(reflect.Type)).Interface()
		err := json.Unmarshal(cmdData, &instance)
		if err != nil {
			return nil, err
		}
		return instance.(model.IEntity), nil
	}
	return nil, fmt.Errorf("DecJsonCmd Not Exists %v", sCmd)
}

//OnRegistCmdEntity  注册cmd实体
func (b bizBase) OnRegistCmdEntity() {
	b.RegistCmd(fmt.Sprintf("%d", model.J平台查岗请求), reflect.TypeOf(down.DOWN_PLATFORM_MSG_INFO_REQ{}))
}
func (b bizBase) RegistCmd(sCmd string, t reflect.Type) {
	b.CmdType.Set(sCmd, t)
}

//分发接口
type IBiz interface {
	SendEntity(obj interface{}) error
	ReceiveEntity() (interface{}, error)
	Start(ctx context.Context) error
}

/*BizDBTdf 数据库分发*/
type BizDBTdf struct {
	bizBase
	TaskRec       *global.TaskManager
	TaskSend      *global.TaskManager
	cmdLRTime     time.Time
	SleepInterval time.Duration //睡眠间隔
}

//Start 启动
func (b *BizDBTdf) Start(ctx context.Context) error {
	b.bizBase.OnRegistCmdEntity() //注册json 反序列化列表
	b.TaskRec.Start(ctx)
	b.TaskSend.Start(ctx)
	go b.OnWork(ctx)
	return nil
}
func (b *BizDBTdf) OnWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		b.ReceiveEntity() //接收cmd
		time.Sleep(b.SleepInterval)
	}
}
func (b *BizDBTdf) Stop() {
	b.TaskRec.Stop()
	b.TaskSend.Stop()
}

//SendEntity 分发数据
func (b *BizDBTdf) SendEntity(obj interface{}) error {
	return nil
}
func (b *BizDBTdf) onSendEntity(obj interface{}) {

}

//ReceiveEntity 接收数据
func (b *BizDBTdf) ReceiveEntity() (interface{}, error) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("BizDBTdf ReceiveEntity panic recover! p: %v", p)
		}
	}()
	stmt, _ := global.DBInstance().Prepare(`select SimNum,UserId,CmdCode,ParamContent,SendTime,CmdCodeFlag from nat_CmdParamEx where SendTime >?  order by SendTime desc`)
	rows, err := stmt.Query(b.cmdLRTime.Format("2006-01-02 15:04:05"))
	if err != nil {
		global.Error("ReceiveEntity Panic.", err)
		return nil, err
	}
	index := 0
	for rows.Next() {
		cmdParam := &cmdParamEx{}
		err = rows.Scan(&cmdParam.SimNum, &cmdParam.UserID, &cmdParam.CmdCode, &cmdParam.ParamContent, &cmdParam.SendTime, &cmdParam.CmdCodeFlag)
		if err != nil {
			global.Error(global.F(global.Biz, global.DOWNDATA, "ReceiveEntity Scan.%v"), err)
			continue
		}
		if index == 0 {
			b.cmdLRTime = cmdParam.SendTime.Add(1 * time.Second)
		}
		entity, err := b.DecCmd("json", cmdParam.CmdCode, []byte(cmdParam.ParamContent))
		if err != nil || entity == nil {
			global.Error(global.F(global.Biz, global.DOWNDATA, "ReceiveEntity DecCmd.%v"), err)
			continue
		}
		global.Debug(global.F(global.Biz, global.DOWNDATA, "Send Cmd %v:%v SendTime:%v"), cmdParam.CmdCode, cmdParam.ParamContent, cmdParam.SendTime.Format("2006-01-02 15:04:05"))
		b.TaskRec.Enqueue(entity)
		index++
	}
	global.Debug(global.F(global.Biz, global.DOWNDATA, "DB GetData Row: 时间:%v Num:%v"), b.cmdLRTime.Format("2006-01-02 15:04:05"), index)
	return nil, nil
}
func (b *BizDBTdf) onReceiveEntity(obj interface{}) {
	if obj == nil {
		return
	}
	if global.DownHandler == nil {
		global.Error(global.F(global.Biz, global.DOWNDATA, "global.DownHandler is nil"))
		return
	}
	rsp, err := global.DownHandler.DownData(obj.(model.IEntity)) //下发指令
	if err != nil {
		global.Error(global.F(global.Biz, global.DOWNDATA, "Biz OnReceiveEntity Error->%v"), err)
		return
	}
	if rsp != nil {
		b.TaskSend.Enqueue(rsp)
	}
}

//NewBizDBTdf 新建一个BizDBTdf
func NewBizDBTdf() *BizDBTdf {
	biz := &BizDBTdf{}
	biz.bizBase = newbizBase()
	biz.TaskRec = global.NewTaskManager("BizDBTdf.Rec.Task", 1000, 1, biz.onReceiveEntity)
	biz.TaskSend = global.NewTaskManager("BizDBTdf.Send.Task", 1000, 1, biz.onSendEntity)
	biz.cmdLRTime = time.Now()
	biz.SleepInterval = time.Millisecond * 5000
	return biz
}

//cmdParamEx  下行指令参数
type cmdParamEx struct {
	SimNum       string
	UserID       string
	CmdCode      string
	ParamContent string
	SendTime     time.Time
	CmdCodeFlag  string
}
