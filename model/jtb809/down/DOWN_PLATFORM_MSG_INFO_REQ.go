package down

//DOWN_PLATFORM_MSG_INFO_REQ 下发平台间报文请求 0x9302
type DOWN_PLATFORM_MSG_INFO_REQ struct{
    model.EntityBase
    //下发报文对象类型
    OBJECT_TYPE byte
    //下发报文对象的ID
    OBJECT_ID  string
    //信息ID
    INFO_ID    uint32
    //消息内容
    INFO_CONTENT string
}

func (e *DOWN_PLATFORM_MSG_INFO_REQ) GetMsgId() interface{} {
	return e.MsgId
}
func (e *DOWN_PLATFORM_MSG_INFO_REQ) GetEntityBase() *EntityBase {
	return e
}

func (e *DOWN_PLATFORM_MSG_INFO_REQ) GetDBSql() string {
	return ""
}