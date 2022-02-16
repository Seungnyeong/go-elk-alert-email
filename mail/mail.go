package mail

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
)

const (
	mailDial string = "wemakeprice-com.mail.protection.outlook.com:25"
	fromAddress string = "wkms@wemakeprice.com"
	mimeString string = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
)

var (
	toAddress string = "seungnyeong@wemakeprice.com"
)

func MailForAdmin(i interface{}) bool {
	fmt.Println("mail to admin")
	if true {
		fmt.Println(i)
	}

	return true
}

func SendMail(){
	
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
	subject := "Subject: Test email from Go!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	body := "<html><body><h1>Hello World!</h1></body></html>"
	buf := bytes.NewBufferString(subject + mime + body)
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
	
}