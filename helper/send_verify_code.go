package helper


import (
	"net/smtp"
	"os"
	"sync"

	emailPKG "github.com/jordan-wright/email"
)


var once sync.Once
var Smtp *SMTP


type SMTP struct {}


func NewEmail() *SMTP {
	once.Do(func() {
			Smtp = &SMTP{}
		})

	return Smtp
}


func (S *SMTP) Send(toUserEmail string, code string) error {
	
	e := emailPKG.NewEmail()

	e.From = os.Getenv("VERIFYCODE_FROM")
	e.To = []string{toUserEmail}
	e.Subject = os.Getenv("EMAIL_SUBJECT")
	e.Text = []byte(code)

	err := e.Send(
		"smtp.qq.com:25",
		smtp.PlainAuth(
			"",
			os.Getenv("VERIFYCODE_FROM"),  // 服务器邮箱账号
			os.Getenv("VERIFYCODE_QQEmailAuthCode"),  // 授权码
			"smtp.qq.com",
		),
	)

	return err
}