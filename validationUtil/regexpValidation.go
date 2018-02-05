package validationUtil

import "regexp"

var (
	// 身份证号
	re_idno = regexp.MustCompile("^[1-9]{1}[0-9]{14}$|^[1-9]{1}[0-9]{16}([0-9]|[xX])$")
	// 邮箱
	re_email = regexp.MustCompile(`^[0-9A-Za-z][\.-_0-9A-Za-z]*@[0-9A-Za-z]+(\.[0-9A-Za-z]+)+$`)

	// 中国手机号验证
	re_chinesePhoneNum = regexp.MustCompile(`^1(3[0-9]|4[57]|5[0-35-9]|7[0135678]|8[0-9])\d{8}$`)
)

// 验证身份证号码
// idno:身份证号
// 返回值:
// bool: true:成功 false:不合法
func IsValideIdno(idno string) bool {
	if re_idno.MatchString(idno) {
		return true
	}

	return false
}

// 验证邮箱
// email:邮箱
// 返回值:
// bool: true:成功 false:不合法
func IsValideEmail(email string) bool {
	if re_email.MatchString(email) {
		return true
	}

	return false
}

// 验证中国的手机号码
// phoneNum:待验证的手机号码
// 返回值:
// bool: true:成功 false:不合法
func IsValideChinesePhoneNum(phoneNum string) bool {
	if re_chinesePhoneNum.MatchString(phoneNum) {
		return true
	}

	return false
}
