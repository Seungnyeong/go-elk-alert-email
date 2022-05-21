package mail

import (
	"fmt"
)

const (
	mimeString string = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	subject    string = "Subject: [중요] WKMS health check alert \n"
)

func SendMail(data interface{}, m chan<- bool) {
	check := true
	html := string(MakeTemplate(data))
	//users, err := service.NewUserRepository().FindAdminUser()
	fmt.Println(html)
	// if err != nil {
	// 	utils.CheckError(err)
	// }
	// buf := bytes.NewBufferString(subject + mimeString + html)

	// c, err := smtp.Dial(config.P.Mail.Host)
	// if err != nil {
	// 	log.Panic("Error", err)
	// 	check = false
	// }

	// defer c.Quit()

	// if err := c.Mail(config.P.Mail.From); err != nil {
	// 	log.Panic(err)
	// 	check = false
	// }

	// for _, user := range users {
	// 	if err := c.Rcpt(user.Email); err != nil {
	// 		log.Panic(err)
	// 		check = false
	// 	}
	// }

	// wc, err := c.Data()

	// if err != nil {
	// 	log.Panic(err)
	// 	check = false
	// }
	// defer wc.Close()

	// if _, err = buf.WriteTo(wc); err != nil {
	// 	log.Panic(err)
	// 	check = false
	// }
	m <- check

}
