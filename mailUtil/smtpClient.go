package mailUtil

// 定义SMTPClient接口
type SMTPClient interface {
	// 设置服务器
	SetServer(host string, port int, isSSL bool)

	// 设置发件箱
	SetSender(name, address, password string)

	//发送邮件:
	//  mailTo 接收方列表
	//  subject 主题
	//  body 正文
	//  isHtmlBody 正文是否html格式
	//  attachFiles 附件
	SendMail(mailTo []string, subject, body string, isHtmlBody bool, attachFiles []string) error
}
