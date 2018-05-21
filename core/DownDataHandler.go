package core

import (
	"github.com/tiptok/gonat/global"
	"github.com/tiptok/gonat/model"
)

//DownHandler 下行处理
type DownHandler struct {
	DownServer global.IDownData /*指令下发服务*/
}

//DownData  指令下行
func (d *DownHandler) DownData(rcv model.IEntity) (model.IEntity, error) {
	return d.DownServer.DownData(rcv)
}

//NewDownHandler 新建一个DownHandler
func NewDownHandler(down global.IDownData) *DownHandler {
	return &DownHandler{
		DownServer: down,
	}
}
