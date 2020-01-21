package qcloud

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"time"

	"github.com/Jordanzuo/goutil/mathUtil"
	"github.com/Jordanzuo/goutil/webUtil"
)

const (
	VOICE_CAPTCHA_URL               = "https://cloud.tim.qq.com/v5/tlsvoicesvr/sendcvoice"
	VOICE_NOTIFICATION_URL          = "https://cloud.tim.qq.com/v5/tlsvoicesvr/sendvoiceprompt"
	VOICE_TEMPLATE_NOTIFICATION_URL = "https://cloud.tim.qq.com/v5/tlsvoicesvr/sendtvoice"
)

// calculate sign-string for phone numbers
func calcSig(appKey string, rand int, timeStamp int64, mobile string) string {
	sum := sha256.Sum256([]byte(fmt.Sprintf("appkey=%s&random=%d&time=%d&mobile=%s", appKey, rand, timeStamp, mobile)))
	return hex.EncodeToString(sum[:])
}

// do the http request and parse the response
func request(url string, data map[string]interface{}, appId string, rand int) (success bool, err error) {
	url = fmt.Sprintf("%s?sdkappid=%s&random=%d", url, appId, rand)
	fmt.Printf("url:%s\n", url)

	contentBytes, err := json.Marshal(data)
	if err != nil {
		return
	}

	fmt.Printf("data:%v\n", string(contentBytes))

	_, retBytes, err := webUtil.PostByteData2(url, contentBytes, nil, nil)
	if err != nil {
		return
	}

	var responseObj *response
	err = json.Unmarshal(retBytes, &responseObj)
	if err != nil {
		return
	}
	if responseObj.Result != 0 {
		err = fmt.Errorf(responseObj.ErrMsg)
		return
	}

	success = true
	return
}

func SendVoiceCaptcha(appId, appKey, nation, mobile, captcha string, playTimes int) (success bool, err error) {
	rand := mathUtil.GetRand().GetRandRangeInt(100000, 999999)
	timeStamp := time.Now().Unix()

	data := make(map[string]interface{})
	data["playtimes"] = playTimes
	data["sig"] = calcSig(appKey, rand, timeStamp, mobile)
	data["tel"] = newTelField(nation, mobile)
	data["time"] = timeStamp
	data["ext"] = ""
	// dedicated param
	data["msg"] = captcha

	success, err = request(VOICE_CAPTCHA_URL, data, appId, rand)
	return
}

func SendVoiceNotification(appId, appKey, nation, mobile, prompt string, playTimes int) (success bool, err error) {
	rand := mathUtil.GetRand().GetRandRangeInt(100000, 999999)
	timeStamp := time.Now().Unix()
	promptType := 2

	data := make(map[string]interface{})
	data["playtimes"] = playTimes
	data["sig"] = calcSig(appKey, rand, timeStamp, mobile)
	data["tel"] = newTelField(nation, mobile)
	data["time"] = timeStamp
	data["ext"] = ""
	// dedicated param
	data["promptfile"] = prompt
	data["prompttype"] = promptType

	success, err = request(VOICE_NOTIFICATION_URL, data, appId, rand)
	return
}

func SendVoiceTemplateNotification(appId, appKey, nation, mobile string, templateId int, params []string, playTimes int) (success bool, err error) {
	rand := mathUtil.GetRand().GetRandRangeInt(100000, 999999)
	timeStamp := time.Now().Unix()

	data := make(map[string]interface{})
	data["playtimes"] = playTimes
	data["sig"] = calcSig(appKey, rand, timeStamp, mobile)
	data["tel"] = newTelField(nation, mobile)
	data["time"] = timeStamp
	data["ext"] = ""
	// dedicated param
	data["tpl_id"] = templateId
	data["params"] = params

	success, err = request(VOICE_TEMPLATE_NOTIFICATION_URL, data, appId, rand)
	return
}
