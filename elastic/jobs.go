package elastic

import (
	"errors"
	"fmt"
	"test/config"
	"test/keyinfo/service"
	"test/mail"
	"test/utils"
)

var ErrCannotExcute = errors.New("cannot excute job for cron")

func checkDowncount(key string) bool {
	down := false
	
	if (is.server[key].Downcount % 10 == 0 && is.server[key].Status == "down") {
		down = true
		is.server[key].UpdateIntanceDownCount(0)
	}

	if (is.server[key].Downcount > 0 && is.server[key].Status == "up") {
		is.server[key].UpdateIntanceDownCount(0)
	}

	return down
}

func updateServerInfo(name string, c chan<- bool)  {
	query := MakeServerMonitoringQuery(name)
	response, err := SearchRestAPIResult(es.Client, &query, "wmp-wkms-health-*")
	result := ParsingInstance(response)
	key := fmt.Sprintf("%s:%s", result.Ip, result.Port)
	utils.CheckError(err)
	check , _ := GetInstance(key, is)

	if check == nil {
		is.AddInstance(result)		
	} else {
		i, err := GetInstance(fmt.Sprintf("%s:%s", result.Ip, result.Port), is)
		utils.CheckError(err)
		i.UpdateIntance(result.Status, utils.RFCtoKST(result.Timestamp))
	}

	c <- checkDowncount(key)
	 
}

func Job(monitorId []string) error {
	if !checkSIEMStatus() {
		errMsg := "cannot connect wmp-siem"
		return  errors.New(errMsg)
	}
	c := make(chan bool)
	
	for _, name := range monitorId {
		go updateServerInfo(name, c)
	}

	if ok :=  <- c; ok {
		users , err := service.NewUserRepository().FindAdminUser()
		utils.CheckError(err)
		
		if err != nil {
			mail.SendMail(config.P.Mail.To, string(MakeTemplate()))
			return err
		}

		for _, user := range users {
			go mail.SendMail(user.Email, string(MakeTemplate()))
		}
		
	}

	return nil
}