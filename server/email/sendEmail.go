package email

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

// Send 发送邮件
// to(string) 发送目标
// subject 主题
// bodyType 正文类型
// body 正文
func Send(to, subject string, body []byte) error {
	e := email.NewEmail()
	e.From = "hardware@zstuca.club"
	e.To = []string{to}
	e.Subject = subject
	e.HTML = body
	return e.Send("imaphz.qiye.163.com:25", smtp.PlainAuth("", e.From, "NqH9zsFgBCDeVyv2", "imaphz.qiye.163.com"))
}
