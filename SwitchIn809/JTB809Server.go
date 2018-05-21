package SwitchIn809

import (
	"log"

	"github.com/tiptok/gonat/global"
	"github.com/tiptok/gonat/model"
	"github.com/tiptok/gotransfer/comm"
	"github.com/tiptok/gotransfer/conn"
)

type Tcp809Server struct {
	Server             conn.TcpServer
	SubList            *comm.DataContext
	SubClientExpireChk *comm.TimerWork
}

func (svr *Tcp809Server) Start() bool {
	//启动tcp服务
	go func() {
		svr.Server.NewTcpServer(global.Param.ServerPort, 500, 500)
		svr.Server.Config.IsParsePartMsg = true //进行分包
		svr.Server.P = &protocol809{}
		svr.SubList = &comm.DataContext{}

		svr.SubClientExpireChk = comm.NewTimerWork()
		taskChkExpire := &comm.Task{
			Interval: 60,
			TaskId:   "ClientExpireCheck",
			Work:     svr.CheckClientExpire,
		}
		svr.SubClientExpireChk.RegistTask(taskChkExpire)

		// taskChkOnlineBuffer := &comm.Task{
		// 	Interval: 60,
		// 	TaskId:   "OnlineExpireCheck",
		// 	Work:     global.OnlineBuffer.ExClean,
		// }

		svr.Server.Start(svr)
		svr.SubClientExpireChk.Start()
	}()
	return true
}

func (svr *Tcp809Server) Stop() bool {
	svr.SubClientExpireChk.Stop()
	return true
}

//CheckClientExpire 检查从链路是否连接超时(无数据交互)
func (svr *Tcp809Server) CheckClientExpire(val interface{}) {
	defer func() {
		if p := recover(); p != nil {
			global.Error("SubClisCacheLoader Load Recover:%v", p)
		}
	}()
	var key string
	svr.SubList.PurgeWithFunc(key, 300, svr.OnRemvoe)
}

func (svr *Tcp809Server) OnRemvoe(k interface{}, val interface{}) {
	defer func() {
		if p := recover(); p != nil {
			global.Error("SubClisCacheLoader OnRemvoe Recover:%v", p)
		}
	}()
	subCli := val.(*TcpSubClient)
	if subCli != nil {
		global.Info("从链路 %v 执行超时处理", subCli.AccessCode)
		subCli.Stop() //超时关闭  做从链路关闭操作
	}
}

//DownData 下发指令
func (trans *Tcp809Server) DownData(rcv model.IEntity) (model.IEntity, error) {
	defer func() {
		if p := recover(); p != nil {
			log.Printf("SvrHander809 DownData panic recover! p: %v", p)
		}
	}()
	var err error
	baseEntity := rcv.GetEntityBase()
	if baseEntity == nil {
		return nil, nil
	}
	if val, ok := trans.SubList.GetOk(baseEntity.AccessCode); ok {
		subcli := val.(*TcpSubClient)
		err = subcli.DownCmd(rcv)
	}
	return nil, err
}
