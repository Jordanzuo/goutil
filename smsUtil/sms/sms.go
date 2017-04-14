package sms

import (
    "github.com/Jordanzuo/goutil/smsUtil/dataSender"
)


type Sms interface {
    // a sms can be used as dataSender.Requester
    dataSender.Requester

    // 用于获取准确的返回数据
    GetResponse() interface{}
}
