package smsUtil

import (
	"testing"

	"github.com/Jordanzuo/goutil/smsUtil/qcloud"
	"github.com/Jordanzuo/goutil/smsUtil/smsSenders/defaultSmsSender"
)

func TestQCloudMsg(t *testing.T) {
	q := qcloud.New("123", "234")
	_, err := q.MsgSms([]string{"86"},
		[]string{"151"},
		"msg",
		"extend",
		"ext")
	if err != nil {
		t.Errorf("qcloud.NewSms (type:msg) error: %v", err)
		return
	}

	_, err = q.TmplSms([]string{"86"},
		[]string{"151"},
		1, []string{"paramenter"},
		"sign", "extend", "ext")
	if err != nil {
		t.Errorf("qcloud.NewSms (type:tmpl) error: %v", err)
		return
	}
}

func TestQCloudSend(t *testing.T) {
	q := qcloud.New("123123123123", "123123123123123213")
	sms, err := q.TmplSms([]string{"86"},
		[]string{"10000000000"},
		14176, []string{"9527"},
		"hello", "", "hello-world!")
	if err != nil {
		t.Errorf("qcloud.NewSms (type:tmpl) error: %v", err)
		return
	}

	sender := defaultSmsSender.New()

	if ok, err := sender.Send(sms); !ok {
		if err != nil {
			t.Errorf("qcloud.Send error: %v", err)
			return
		} else if !ok {
			rspn := sms.GetResponse().(*qcloud.QCloudResponse)
			t.Errorf("qcloud return error: %v", rspn)
		}
	}
}
