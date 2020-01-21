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
	MULTI_SMS_URL  = "https://yun.tim.qq.com/v5/tlssmssvr/sendmultisms2"
)

// 创建msg类型短信(直接发送短信内容)
// 参数:
//      appid
//      appkey
//      nation      []string    国家代码 eg: 86
//      numbers     []string    电话号码
//      msg         string      短信内容
//      extend      string      通道扩展码，可选字段，默认没有开通(需要填空串)
//      ext         string      用户的session内容，腾讯server回包中会原样返回，可选字段，不需要就填空
//
//      nation、numbers 需要一一对应 (nation[0]将与numbers[0]组合)
//      msg 必须与已通过审核的模板匹配
func NewMsgSms(appid, appkey string,
	nation, numbers []string,
	msg, extend, ext string) (sms.Sms, error) {

	return generateSms(appid, appkey,
		nation, numbers,
		extend, ext,
		newMsgSmsField(0, msg),
		nil)
}

// 创建tmpl类型短信(通过模板发送短信)
// 参数:
//      appid
//      appkey
//      nation      []string    国家代码 eg: 86
//      numbers     []string    电话号码
//      tmplID      int         模板ID
//      params      []string    模板参数
//      sign        string      短信签名(前缀)
//      extend      string      通道扩展码，可选字段，默认没有开通(需要填空串)
//      ext         string      用户的session内容，腾讯server回包中会原样返回，可选字段，不需要就填空
//
//      nation、numbers 需要一一对应 (nation[0]将与numbers[0]组合)
func NewTmplSms(appid, appkey string,
	nation, numbers []string,
	tmplID int, params []string,
	sign, extend, ext string) (sms.Sms, error) {

	return generateSms(appid, appkey,
		nation, numbers,
		extend, ext,
		nil,
		newTmplSmsField(sign, tmplID, params))

}

// create qcloud sms instance
func generateSms(appid, appkey string,
	nation, numbers []string,
	extend, ext string,
	msg *msgSmsField,
	tmpl *tmplSmsField) (s sms.Sms, err error) {

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
	rnd := mathUtil.GetRand().GetRandRangeInt(100000, 999999)
	timestap := time.Now().Unix()
	sig := calcSig(appkey, rnd, timestap, numbers)

	// 生成公共字段
	comField := newCommonField(sig, timestap, extend, ext, generateTelField(nation, numbers))

	// 生成smsData
	data := newSmsData(comField, msg, tmpl)

	// 生成qcloudsms
	s = newSms(url, data, rnd, appid)

	return
}

// calculate sign-string for phone numbers
func calcSig(appkey string, rnd int, tm int64, numbers []string) string {
	mobile := strings.Join(numbers, ",")

	sum := sha256.Sum256([]byte(fmt.Sprintf("appkey=%v&random=%v&time=%v&mobile=%v",
		appkey, rnd, tm, mobile)))

	return hex.EncodeToString(sum[:])
}

// 验证地区代码与手机号数量是否匹配
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
	res := make([](*telField), len(nation))

	for i := 0; i < len(nation); i++ {
		res[i] = newTelField(nation[i], numbers[i])
	}

	if len(res) == 1 {
		return res[0]
	}

	return res
}
