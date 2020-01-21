package mailUtil

import (
	"testing"
)

func TestSendMail(t *testing.T) {
	svr := SimpleSMTPClient("smtp.exmail.qq.com", 465, true, "name", "service@moqikaka.com", "Sv123456")
	err := svr.SendMail([]string{"164760769@qq.com"},
		"邮件发送测试",
		"<h1>这是邮件正文</h1>",
		true,
		[]string{
			"./doc.go",
			"./simpleClient_test.go",
		})
	if err != nil {
		t.Error(err)
	}
}
