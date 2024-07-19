package helper

import (
	"blog_backend/models"
	"bytes"
	"fmt"
	"github.com/jordan-wright/email"
	"html/template"
	"net/smtp"
)

type SendEmailParams struct {
	To       string
	Subject  string
	Code     string
	UserInfo *models.User
}

func SendEmail(params SendEmailParams) error {
	tmpl, err := template.ParseFiles("./template/email.html")
	if err != nil {
		return err
	}

	body := new(bytes.Buffer)

	err = tmpl.Execute(body, map[string]interface{}{
		"Code":     params.Code,
		"UserName": params.UserInfo.Username,
		"Avatar":   fmt.Sprintf("http://minio.boyyang.cn/boyyang/%s", params.UserInfo.Avatar),
		"Cover":    fmt.Sprintf("http://minio.boyyang.cn/boyyang/%s", params.UserInfo.Cover),
		"Subject":  params.Subject,
	})

	if err != nil {
		return err
	}

	e := &email.Email{
		From:    "1761617270@qq.com",
		To:      []string{params.To},
		Subject: params.Subject,
		HTML:    body.Bytes(),
	}

	err = e.Send("smtp.qq.com:587", smtp.PlainAuth("", "1761617270@qq.com", "oxixpvgbskpwhfdi", "smtp.qq.com"))
	if err != nil {
		return err
	}

	return nil
}
