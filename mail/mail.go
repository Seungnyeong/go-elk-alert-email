package mail

import (
	"bytes"
	"log"
	"net/smtp"
)

const (
	mailDial string = "wemakeprice-com.mail.protection.outlook.com:25"
	fromAddress string = "wkms@wemakeprice.com"
	mimeString string = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject string = "Subject: [중요] WKMS health check alert \n"
)

var (
	toAddress string = "seungnyeong@wemakeprice.com"
)


func SendMail(html string){
	buf := bytes.NewBufferString(subject + mimeString + html)
	
	c, err := smtp.Dial(mailDial)
	if err != nil {
		log.Fatal("Error", err)
	}
	
	defer c.Quit()
	
	if err := c.Mail(fromAddress); err != nil {
		log.Fatal(err)
	}

	if err := c.Rcpt(toAddress); err != nil {
		log.Fatal(err)
	}

	wc , err := c.Data()
	if err != nil {
		log.Fatal(err)
	}
	defer wc.Close()
	
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
	
}