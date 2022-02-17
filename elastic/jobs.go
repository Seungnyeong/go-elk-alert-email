package elastic

import (
	"errors"
	"fmt"
	"log"

	"test/mail"
	"test/utils"
	"time"

	"github.com/go-co-op/gocron"
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
	GetSingleton()
	es, _ := ElasticConnection()
	res, err := es.Info()
	utils.CheckError(err)
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	
	for _, name := range monitorId {
		result := elsticResult(es, name)
		if !FindInstance(fmt.Sprintf("%s:%s", result.Ip, result.Port)) {
			is.AddInstance(result)		
		} else {
			i, err := is.GetInstance(fmt.Sprintf("%s:%s", result.Ip, result.Port))
			utils.CheckError(err)
			i.UpdateIntance(result.Status, utils.RFCtoKST(result.Timestamp))
			
		}
	}
	
	if CheckDowncount() {
		mail.SendMail(string(MakeTemplate()))
	}	

	if err != nil {
		err = ErrCannotExcute
	}

	return err
}

func CronJob(monitorId []string) error {
	s := gocron.NewScheduler(time.UTC)
	_, err := s.SingletonMode().Every(5).Second().Do(func ()  {
		Job(monitorId)
	})
	if err != nil {
		err = ErrCannotExcute
	}
	s.StartAsync()	
	return err
}

