package helper

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

func SendEmail(to string, text string) error {
	e := &email.Email{
		From:    "1761617270@qq.com",
		To:      []string{to},
		Subject: "boyyang 博客账号注册",
		Text:    []byte(text),
	}

	//oxixpvgbskpwhfdi
	err := e.Send("smtp.qq.com:587", smtp.PlainAuth("", "1761617270@qq.com", "oxixpvgbskpwhfdi", "smtp.qq.com"))
	if err != nil {
		return err
	}

	return nil
}
