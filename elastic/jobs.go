package elastic

import (
	"errors"
	"fmt"

	"test/mail"
	"test/utils"
)

var ErrCannotExcute = errors.New("cannot excute job for cron")
var c = make(chan *Instance)

func checkDowncount(key string) bool {
	down := false

	if is.server[key].Downcount%10 == 0 && is.server[key].Status == "down" {
		down = true
		is.server[key].UpdateIntanceDownCount(0)
	}

	if is.server[key].Downcount > 0 && is.server[key].Status == "up" {
		is.server[key].UpdateIntanceDownCount(0)
		is.server[key].Mailed = false
	}

	return down
}

func updateServerInfo(name string, c chan<- *Instance) {
	query := MakeServerMonitoringQuery(name)
	response, err := SearchRestAPIResult(es.Client, &query, "wmp-wkms-health-*")
	result := ParsingInstance(response)
	key := fmt.Sprintf("%s:%s", result.Ip, result.Port)
	utils.CheckError(err)
	i, _ := GetInstance(key, is)

	if i == nil {
		i = is.AddInstance(result)
	} else {
		utils.CheckError(err)
		i.UpdateIntance(result.Status, utils.RFCtoKST(result.Timestamp))
	}
	checkDowncount(key)
	c <- i
}

func Job(monitorId []string) error {
	if !checkSIEMStatus() {
		errMsg := "cannot connect wmp-siem"
		return errors.New(errMsg)
	}
	for _, name := range monitorId {
		go updateServerInfo(name, c)

		if i := <-c; i.Status == "down" && i.Downcount%10 == 0 {
			if !i.Mailed {
				t := make(chan bool)
				go mail.SendMail(i, t)
				if <-t {
					i.UpdateMailed()
				}
			}
		}
	}
	return nil
}
