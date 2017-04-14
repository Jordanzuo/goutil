package qcloud

import (
    "github.com/Jordanzuo/goutil/logUtil"
    "github.com/Jordanzuo/goutil/smsUtil/sms"
)

type qcloud struct {
    appid  string
    appkey string
}

// 创建msg类型短信(直接发送短信内容)
func (this *qcloud) MsgSms(nation, numbers []string,
                        msg, extend, ext string) (s sms.Sms, err error) {
    s, err = generateSms(this.appid, this.appkey,
                        nation, numbers,
                        extend, ext,
                        &msgSmsField{Type: 0, Msg: msg},
                        nil)
    if err != nil {
        logUtil.NormalLog(err.Error(), logUtil.Error)
    }

    return
}

// 创建tmpl类型短信(通过模板发送短信)
func (this *qcloud) TmplSms(nation, numbers []string,
                        tmplID int, params []string,
                        sign, extend, ext string) (s sms.Sms, err error) {

    s, err = generateSms(this.appid, this.appkey,
                        nation, numbers,
                        extend, ext,
                        nil,
                        &tmplSmsField{Sign: sign, Tpl_id: tmplID, Params:params})
    if err != nil {
        logUtil.NormalLog(err.Error(), logUtil.Error)
    }

    return
}

func New(appid, appkey string) *qcloud {
    return &qcloud{appid: appid, appkey: appkey}
}
