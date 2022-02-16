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
	buf := bytes.NewBufferString("subject: this is Subject\n\n <h1>this is the body</h1>\n")
	if _, err = buf.WriteTo(wc); err != nil {
		log.Fatal(err)
	}
	
}