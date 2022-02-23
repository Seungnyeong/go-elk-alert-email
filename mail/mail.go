package mail

import (
	"fmt"
)

const (
	mimeString string = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject    string = "Subject: [중요] WKMS health check alert \n"
)

func SendMail(user_mail string, html string) {

	//buf := bytes.NewBufferString(subject + mimeString + html)
	fmt.Println(user_mail)
	// c, err := smtp.Dial(config.P.Mail.Host)
	// if err != nil {
	// 	log.Panic("Error", err)
	// }

	// defer c.Quit()

	// if err := c.Mail(config.P.Mail.From); err != nil {
	// 	log.Panic(err)
	// }

	// if err := c.Rcpt(user_mail); err != nil {
	// 	log.Panic(err)
	// }

	// wc , err := c.Data()
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer wc.Close()

	// if _, err = buf.WriteTo(wc); err != nil {
	// 	log.Panic(err)
	// }

}
