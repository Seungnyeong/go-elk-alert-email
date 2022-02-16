package mail

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	mailDial string = "wemakeprice-com.mail.protection.outlook.com:25"
	fromAddress string = "test@wemakeprice.com"
	toAddress string = "seungnyeong@wemakeprice.com"
)

func MailForAdmin(i interface{}) bool {
	fmt.Println("mail to admin")
	if true {
		fmt.Println(i)
	}

	return true
}

func Start(){
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

	_, err = fmt.Fprintf(wc, "This is the email body")
	if err != nil {
		log.Fatal(err)
	}

	err = wc.Close()
	if err != nil {
		log.Fatal(err)
	}

}