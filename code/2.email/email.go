package email

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

//代码教程 https://www.topgoer.cn/docs/goday/goday-1crg220gnel8q
func Send(addr, subject, text string) error {
	// qq邮箱服务器有问题 配置不好
	e := email.NewEmail()
	e.From = "d136169<d136169@163.com>" // 发送方
	e.To = []string{addr}               // 接收方
	e.Subject = subject                 // 标题
	e.Text = []byte(text)               // 正文
	err := e.Send("smtp.163.com:25", smtp.PlainAuth("", "d136169@163.com", "EGRULHVXOZYGDEVF", "smtp.163.com"))
	return err
}
