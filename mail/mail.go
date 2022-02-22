package mail

import (
	"bytes"
	"fmt"
	"log"
	"net/smtp"
	"test/config"
	"test/keyinfo/domain"
	"test/keyinfo/service"
)

const (
	mimeString string = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject string = "Subject: [중요] WKMS health check alert \n"
)


func SendMail(html string){
	user , err := service.NewUserRepository().FindUser("seungnyeong")
	if err != nil {
		user = domain.User{
			Username: "seungnyeong",
			Email: "seungnyeong@wemakeprice.com",
			IsSuperUser: true,
			IsActive: false,
		}
	}

	fmt.Println("sending email")
	buf := bytes.NewBufferString(subject + mimeString + html)
	
	c, err := smtp.Dial(config.Properties().Mail.Host)
	if err != nil {
		log.Fatal("Error", err)
	}
	
	defer c.Quit()
	
	if err := c.Mail(config.Properties().Mail.From); err != nil {
		log.Fatal(err)
	}

	if err := c.Rcpt(user.Email); err != nil {
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