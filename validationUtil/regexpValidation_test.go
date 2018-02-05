package validationUtil

import (
	"testing"
)

// 身份证测试
func TestIdCard(t *testing.T) {
	idno := "450325197410077393"
	if IsValideIdno(idno) == false {
		t.Error("身份证验证出错:", idno)
		t.Fail()
	}
	idno = "36062219701120774X"
	if IsValideIdno(idno) == false {
		t.Error("身份证验证出错:", idno)
		t.Fail()
	}
	idno = "450325197410071111"
	if IsValideIdno(idno) == false {
		t.Error("身份证验证出错:", idno)
		t.Fail()
	}
	idno = "3123123123"
	if IsValideIdno(idno) == true {
		t.Error("身份证验证出错:", idno)
		t.Fail()
	}
}

// 邮箱测试
func TestMail(t *testing.T) {
	mail := "nihao@qq.com"
	if IsValideEmail(mail) == false {
		t.Error("邮箱验证出错:", mail)
		t.Fail()
	}
	mail = "111@qq.com"
	if IsValideEmail(mail) == false {
		t.Error("邮箱验证出错:", mail)
		t.Fail()
	}
	mail = "111_@qq.com"
	if IsValideEmail(mail) == false {
		t.Error("邮箱验证出错:", mail)
		t.Fail()
	}
}

// 验证中国的手机号
func TestChinesePhone(t *testing.T) {
	phoneNum := "15111111111"
	if IsValideChinesePhoneNum(phoneNum) == false {
		t.Error("手机号验证出错:", phoneNum)
		t.Fail()
	}
	phoneNum = "11111"
	if IsValideChinesePhoneNum(phoneNum) == true {
		t.Error("手机号验证出错:", phoneNum)
		t.Fail()
	}
}
