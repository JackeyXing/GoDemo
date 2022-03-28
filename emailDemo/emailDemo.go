//emial库
//利用qq服务器
package main

import (
	"net/smtp"

	"github.com/jordan-wright/email"
)

func main() {
	e := email.NewEmail()
	e.From = "xxj <2810947907@qq.com>"
	e.To = []string{"2810947907@qq.com"}
	e.Bcc = []string{"2810947907@qq.com"}
	e.Subject = "测试golang email lib"
	e.Text = []byte("以下是内容")
	e.HTML = []byte("<h1>content here</h1>")
	e.Send("smtp.qq.com:587", smtp.PlainAuth("",
		"2810947907@qq.com", "password here", "smtp.qq.com"))

	// //Another Method for Creating an Email
	// e := &email.Email{
	// 	To:      []string{"2810947907@qq.com"},
	// 	From:    "xxj <2810947907@q.com>",
	// 	Subject: "test golang email lib",
	// 	Text:    []byte("以下是内容"),
	// 	HTML:    []byte("<h1>content here</h1>"),
	// 	Headers: textproto.MIMEHeader{},
	// 	Bcc:     []string{"2810947907@qq.com"},
	// }

	// e.Send("smtp.qq.com:7777", smtp.PlainAuth("",
	// 	"2810947907@qq.com", "ahlaiocylratdeca", "smtp.qq.com"))

	// // Creating an Email From an io.Reader
	// e := email.NewEmail()
	// e.AttachFile("text.txt")
	// e.Send("smtp.qq.com:9000", smtp.PlainAuth("",
	// 	"2810947907@qq.com", "ahlaiocylratdeca", "smtp.qq.com"))

	// // A Pool of Reusable Connections
	// var ch <-chan *email.Email
	// p := email.NewPool(
	// 	"smtp.qq.com:587",
	// 	4,
	// 	smtp.PlainAuth("", "2810947907@qq.com", "ahlaiocylratdeca", "smtp.qq.com"),
	// )
	// for i := 0; i < 10; i++ {
	// 	go func() {
	// 		for e := range ch {
	// 			p.Send(e, 10*time.Second)
	// 		}
	// 	}()
	// }

}
