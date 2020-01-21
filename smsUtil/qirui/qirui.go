package qirui

import (
	"fmt"

	"github.com/Jordanzuo/goutil/webUtil"
)

const (
	SEND_MESSAGE_URL = "http://api.qirui.com:7891/mt"
)

func SendMessage(apiKey, apiSecret, mobile, message string) (bool, error) {
	smsUrl := fmt.Sprintf("%s?dc=15&un=%s&pw=%s&da=%s&sm=%s&tf=3&rf=2&rd=1", SEND_MESSAGE_URL, apiKey, apiSecret, mobile, message)
	status, resp, err := webUtil.GetWebData3(smsUrl, "", webUtil.ContentType_Json, nil)
	if err != nil {
		return false, err
	}

	if status != 200 {
		return false, fmt.Errorf("StatusCode is %d", status)
	}

	fmt.Printf("qirui.SendMessage:%s, %s, %s\n", mobile, message, string(resp))
	return true, nil
}
