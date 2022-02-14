package mail

import (
	"fmt"
	"log"
	"net/smtp"
)

func Start(){
	c, err := smtp.Dial("wemakeprice-com.mail.protection.outlook.com:25")
	if err != nil {
		log.Fatal("Error", err)
	}
	
	defer c.Quit()
	
	if err := c.Mail("test@wemakeprice.com"); err != nil {
		log.Fatal(err)
	}

	if err := c.Rcpt("seungnyeong@wemakeprice.com"); err != nil {
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