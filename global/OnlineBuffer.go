package global

import (
	"github.com/tiptok/GoNas/global"
	"github.com/tiptok/gotransfer/comm"
)

type OnlineList struct {
	List *comm.DataContext
}

//Login  上线
func (online *OnlineList) Login(info *VehicleInfo) {
	Info(global.F(global.BUS, global.Cache, " %v 上线车辆 %v %v"), info.Obj, info.PlateNum, info.SimNum)
	online.List.Set(info.SimNum, info)
}

//Logout 下线
func (online *OnlineList) Logout(sKey string) {
	if val, ok := online.List.GetOk(sKey); ok {
		online.List.Delete(sKey)
		online.OnDelete(sKey, val)
	}
}

//OnDelete 关闭连接
func (online *OnlineList) OnDelete(sKey interface{}, val interface{}) {
	info := val.(*VehicleInfo)
	info.SetObj(nil)
	Info(global.F(global.BUS, global.Cache, " %v 下线车辆 %v %v"), info.Obj, info.PlateNum, info.SimNum)
}

//GetOnline 返回在线车辆
func (online *OnlineList) GetOnline(sKey string) *VehicleInfo {
	info := online.List.Get(sKey)
	return info.(*VehicleInfo)
}

//超时清理
func (online *OnlineList) ExClean(param interface{}) {
	var sKey string
	online.List.PurgeWithFunc(sKey, 60*2, online.OnDelete)
}
