package elastic

import (
	"errors"
	"fmt"
	"test/mail"
	"test/utils"
)

var ErrCannotExcute = errors.New("cannot excute job for cron")

func CheckDowncount() bool {
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



func Job(monitorId []string) error {
	var errMsg string
	if !checkSIEMStatus() {
		errMsg = "cannot connect wmp-siem"
	}

	for _, name := range monitorId {
		query := MakeServerMonitoringQuery(name)
		response, err := SearchRestAPIResult(es.Client, &query, "wmp-wkms-health-*")
		result := ParsingInstance(response)

		if err != nil {
			errMsg = err.Error()
		}
		check , _ := GetInstance(fmt.Sprintf("%s:%s", result.Ip, result.Port), is)
		if check == nil {
			is.AddInstance(result)		
		} else {
			i, err := GetInstance(fmt.Sprintf("%s:%s", result.Ip, result.Port), is)
			utils.CheckError(err)
			i.UpdateIntance(result.Status, utils.RFCtoKST(result.Timestamp))
		}
	}
	
	if CheckDowncount() {
		mail.SendMail(string(MakeTemplate()))
	}
	return errors.New(errMsg)
}


