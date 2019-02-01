package down

import (
	"github.com/tiptok/gonat/model"
)

//SendUpgradePackage 下发终端升级包 0x8108
type SendUpgradePackage struct {
	model.EntityBase        //升级类型 0：终端，12：道路运输证IC卡读卡器，52：北斗卫星定位模块
	ManufacturerId   string //制造商ID string 长度5
	SoftwareVersion  string //版本号
	DataPackage      []byte // 升级包大小
	PackageCount     int    //包数
	PackageSN        int    //包序
	UpgradeType      byte   //升级类型
	FileLength       int    //文件大小
}

func (e *SendUpgradePackage) GetMsgId() interface{} {
	return model.JSendUpgradePackage
}
func (e *SendUpgradePackage) GetEntityBase() *model.EntityBase {
	return e.EntityBase.GetEntityBase()
}

func (e *SendUpgradePackage) GetDBSql() string {
	return ""
}
