package tools

import (
	"strings"

	"github.com/go-gomail/gomail"
	"github.com/spf13/viper"
)

// SetEmail .发送邮件
// toers 接收邮箱
// subject 邮件主题
// body 邮件内容
func SetEmail(toers []string, subject, body string) error {

	// 获取邮件信息
	m := gomail.NewMessage()

	for index, tmp := range toers {
		toers[index] = strings.TrimSpace(tmp)
	}

	// 收件人可以有多个，故用此方式
	m.SetHeader("To", toers...)

	// 发件人
	// 第三个参数为发件人别名，如"李大锤"，可以为空（此时则为邮箱名称）
	m.SetAddressHeader("From", viper.GetString("Emails.fromemails"), viper.GetString("Emails.sender"))

	// // 主题
	m.SetHeader("Subject", subject)

	// // 正文
	m.SetBody("text/html", body)
	d := gomail.NewDialer(viper.GetString("Emails.serverhost"), viper.GetInt("Emails.Serverport"), viper.GetString("Emails.fromemails"), viper.GetString("Emails.password"))
	// // 发送
	err := d.DialAndSend(m)
	return err
}
