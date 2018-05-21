package core

import (
	"log"

	"github.com/tiptok/gonat/SwitchIn808"
	"github.com/tiptok/gonat/SwitchIn809"
	"github.com/tiptok/gonat/global"
	"github.com/tiptok/gotransfer/conn"
)

type Host struct {
	NasServer conn.ITcpServer
	CacheMana global.CacheManage
	//上行处理
	//分析处理
	//下行处理
	//分发处理
	//入库处理
}

func (h *Host) Start(protocol string) {
	var init bool = false

	//加载缓存管理
	h.CacheMana = global.CacheManage{}
	init = h.CacheMana.Init()

	switch protocol {
	case "JTB808":
		h.NasServer = &SwitchIn808.Tcp808Server{}
	case "JTB809":
		server809 := &SwitchIn809.Tcp809Server{}
		h.NasServer = server809
		global.DownHandler = NewDownHandler(server809)
	}
	uphandler := &Up808Data{}
	uphandler.BizDB = NewMSDBHandler()
	global.UpHandler = uphandler
	//global.d
	if h.NasServer != nil {
		init = h.NasServer.Start()
	}
	log.Printf("Host Start %s Result:%v", protocol, init)
}
