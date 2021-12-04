package email

import (
	"bytes"
	"crypto/tls"
	"github.com/wenknow/gin-vue-blog/server/global"
	"github.com/wenknow/gin-vue-blog/server/utils/fileopt"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
	"path/filepath"
)

type LoginEmailInfo struct {
	MailTo string
	Url    string
	Code   string
}

func (r LoginEmailInfo) SendCodeEmail() error {

	subject := "欢迎登录文知道-开发者的博客平台"
	html := "code.html"
	return r.sendHtmlEmail(subject, html)

}

func (r LoginEmailInfo) SendRegisterEmail() error {

	subject := "欢迎注册文知道-开发者的博客平台"
	html := "register.html"
	return r.sendHtmlEmail(subject, html)
}

func (r LoginEmailInfo) sendHtmlEmail(subject string, html string) error {

	path := fileopt.GetCurrentAbPathByCaller()
	path = filepath.Join(path, "/resource/email/"+html)
	t := template.New(html)

	var err error
	t, err = t.ParseFiles(path)
	if err != nil {
		log.Println(err)
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, r); err != nil {
		log.Println(err)
	}

	result := tpl.String()

	err = sendMail([]string{r.MailTo}, subject, result)
	return err
}

func sendMail(mailTo []string, subject string, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", m.FormatAddress(global.CONFIG.Email.Sender, global.CONFIG.Email.Name)) //这种方式可以添加别名，即“XX官方”
	//说明：如果是用网易邮箱账号发送，以下方法别名可以是中文，如果是qq企业邮箱，以下方法用中文别名，会报错，需要用上面此方法转码
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正文

	d := gomail.NewDialer(
		global.CONFIG.Email.Host,
		global.CONFIG.Email.Port,
		global.CONFIG.Email.Sender,
		global.CONFIG.Email.Secret,
	)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: global.CONFIG.Email.IsSSL}

	err := d.DialAndSend(m)
	return err

}
