package qcloud

import (
    "crypto/sha256"
    "encoding/hex"
    "errors"
    "fmt"
    "strings"
    "time"
    "github.com/Jordanzuo/goutil/mathUtil"
    "github.com/Jordanzuo/goutil/smsUtil/sms"
)

const (
    SINGLE_SMS_URL = "https://yun.tim.qq.com/v5/tlssmssvr/sendsms"
    MULTI_SMS_URL = "https://yun.tim.qq.com/v5/tlssmssvr/sendmultisms2"
)


// create qcloud sms
// 参数:
//      appid
//      appkey
//      nation      []string    国家代码 eg: 86
//      numbers     []string    电话号码
//      extend      string      通道扩展码，可选字段，默认没有开通(需要填空串)
//      ext         string      用户的session内容，腾讯server回包中会原样返回，可选字段，不需要就填空
//
//      nation、numbers 需要一一对应 (nation[0]将与numbers[0]组合)
//      msg 必须与已通过审核的模板匹配
//
func generateSms(appid, appkey string,
            nation, numbers []string,
            extend, ext string,
            msg *msgSmsField, tmpl *tmplSmsField) (s sms.Sms, err error) {


    err = validateMobile(nation, numbers)
    if err != nil {
        return
    }

    var url string
    // 区分
    if len(nation) > 1 {
        url = MULTI_SMS_URL
    } else {
        url = SINGLE_SMS_URL
    }

    // 生成随机数、时间戳，并计算签名
    rnd      := mathUtil.GetRandRangeInt(100000, 999999)
    timestap := time.Now().Unix()
    sig      := calcSig(appkey, rnd, timestap, numbers)

    // 填充公共字段
    comField := &commonField{Sig: sig, Extend: extend, Ext: ext, Time: timestap,
                            Tel: generateTelField(nation, numbers)}

    // 生成smsData
    data := &smsData{commonField: comField, msgSmsField: msg, tmplSmsField: tmpl}

    // url 额外参数
    url = fmt.Sprintf("%s?sdkappid=%s&random=%d", url, appid, rnd)

    // 填充sms
    s = &qcloudsms{data: data, url: url,}

    return
}


// calculate sign-string for phone numbers
func calcSig(appkey string, rnd int, tm int64, numbers []string) string {
    mobile := strings.Join(numbers, ",")

    sum := sha256.Sum256([]byte(fmt.Sprintf("appkey=%v&random=%v&time=%v&mobile=%v",
                                             appkey, rnd, tm, mobile)))

    return hex.EncodeToString(sum[:])
}

func validateMobile(nation, numbers []string) error {
    if len(nation) != len(numbers) {
        return errors.New("loadReq: nation、numbers 数量不同")
    } else if len(nation) == 0 {
        return errors.New("loadReq: nation、numbers 不能为空")
    }

    return nil
}

// 生成Tel字段
// 当多个号码时，生成[]telField
// 单个号码时，生成telField
func generateTelField(nation, numbers []string) interface{} {
    res := make([]telField, len(nation))

    for i := 0; i < len(nation); i++ {
        res[i] = telField{
            Nationcode: nation[i],
            Mobile: numbers[i],
        }
    }

    if len(res) == 1 {
        return res[0]
    }

    return res
}
