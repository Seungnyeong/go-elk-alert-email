package crons

import (
	"errors"
	"test/elastic"
	"time"

	"github.com/go-co-op/gocron"
)


var ErrCannotFindIDs = errors.New("cannot Find Monitoring ID from you request")

func MonitorInstanceJob(ipv4 string) error {
	elastic.GetSingleton()
	s := gocron.NewScheduler(time.UTC)
	query := elastic.MakeServerGroupQuery(ipv4)
	response, err := elastic.SearchRestAPIResult(elastic.ElasticClient().Client , &query, "wmp-wkms-health-*")
	
	if err != nil {
		return err
	}
	
	motoringIds := elastic.ParsingInstanceId(response)

	if !(len(motoringIds) > 0) {
		return ErrCannotFindIDs
	}

	cr, err := s.SingletonMode().Every(5).Second().Do(func () {
		elastic.Job(motoringIds)
	})

	if err != nil {
		s.Remove(cr)
		s.Stop()
	} else {
		s.StartAsync()	
	}
	 
	return err
}
