package email

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

// Send 发送邮件
// to(string) 发送目标
// emailAddr 邮箱服务器
// emailPort 邮箱服务器端口
// emailPassword 邮箱服务密钥
// subject 主题
// bodyType 正文类型
// body 正文
func Send(from, to, emailAddr, emailPort, emailPassword, subject string, body []byte) error {
	return SendAll(from, []string{to}, emailAddr, emailPort, emailPassword, subject, body)
}

// SendAll 向多个目标发送邮件
// from 发件人
// to([]string) 发送目标们
// emailAddr 邮箱服务器
// emailPort 邮箱服务器端口
// emailPassword 邮箱服务密钥
// subject 主题
// bodyType 正文类型
// body 正文
func SendAll(from string, to []string, emailAddr, emailPort, emailPassword, subject string, body []byte) error {
	e := email.NewEmail()
	e.From = from
	e.To = to
	e.Subject = subject
	e.HTML = body
	return e.Send(emailAddr+":"+emailPort, smtp.PlainAuth("", e.From, emailPassword, emailAddr))
}
