package elastic

import (
	"errors"
	"fmt"

	"test/instances/instance"
	inst "test/instances/service"
	"test/mail"
	"test/utils"
)

var (
	ErrCannotExcute = errors.New("cannot excute job for cron")
	c               = make(chan *instance.Instance)
)

type JobRepository struct {
	elk ElasticAccess
}

func NewJob() *JobRepository {
	return &JobRepository{Client()}
}

func (job JobRepository) updateServerInfo(name string, c chan<- *instance.Instance) {
	query := MakeServerMonitoringQuery(name)
	response, err := job.elk.Search(&query, "wmp-wkms-health-*")
	result := ParsingInstance(response)
	key := fmt.Sprintf("%s:%s", result.Ip, result.Port)
	utils.CheckError(err)
	i, _ := inst.GetInstance(key)

	if i == nil {
		i = inst.NewInstances().AddInstance(result)
	} else {
		utils.CheckError(err)
		i.UpdateIntance(result.Status, utils.RFCtoKST(result.Timestamp))
	}
	inst.NewInstances().UpdateDownCount(key)
	c <- i
}

func (job JobRepository) Job(monitorId []string) error {
	for _, name := range monitorId {
		go job.updateServerInfo(name, c)

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
