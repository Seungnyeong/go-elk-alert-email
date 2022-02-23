package elastic

import (
	"errors"
	"fmt"
	"test/mail"
	"test/utils"
)

var ErrCannotExcute = errors.New("cannot excute job for cron")

func checkDowncount() bool {
	down := false
	for _, server := range is.server {
		if ( server.Downcount % 10 == 0) && server.Status == "down" {
			down = true
			server.UpdateIntanceDownCount(0)
			break;
		} 

		if server.Downcount > 0 && server.Status == "up" {
			server.UpdateIntanceDownCount(0)
		}
	}
	return down
}

func updateServerInfo(name string){
	query := MakeServerMonitoringQuery(name)
	response, err := SearchRestAPIResult(es.Client, &query, "wmp-wkms-health-*")
	result := ParsingInstance(response)
	utils.CheckError(err)
	check , _ := GetInstance(fmt.Sprintf("%s:%s", result.Ip, result.Port), is)

	if check == nil {
		is.AddInstance(result)		
	} else {
		i, err := GetInstance(fmt.Sprintf("%s:%s", result.Ip, result.Port), is)
		utils.CheckError(err)
		i.UpdateIntance(result.Status, utils.RFCtoKST(result.Timestamp))
	}
}

func Job(monitorId []string) error {
	if !checkSIEMStatus() {
		errMsg := "cannot connect wmp-siem"
		return  errors.New(errMsg)
	}

	for _, name := range monitorId {
		go updateServerInfo(name)
	}
	
	if checkDowncount() {
		mail.SendMail(string(MakeTemplate()))
	}

	return nil
}