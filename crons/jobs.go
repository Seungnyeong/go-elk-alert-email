package crons

import (
	"errors"
	"test/elastic"
	"time"

	"github.com/go-co-op/gocron"
)

var ErrCannotFindIDs = errors.New("cannot Find Monitoring ID from you request")
var ErrSIEMConnect = errors.New("cannot Connect SIEM")

func MonitorInstanceJob(ipv4 string) error {
	var err error
	s := gocron.NewScheduler(time.UTC)
	status := elastic.Client().Status()

	if !status {
		return ErrSIEMConnect
	}

	query := elastic.MakeServerGroupQuery(ipv4)
	response, err := elastic.Client().Search(&query, "wmp-wkms-health-*")

	if err != nil {
		return err
	}

	motoringIds := elastic.ParsingInstanceId(response)

	if !(len(motoringIds) > 0) {
		return ErrCannotFindIDs
	}

	cr, err := s.SingletonMode().Every(5).Second().Do(func() {
		err = elastic.NewJob().Job(motoringIds)
	})

	if err != nil {
		s.Remove(cr)
		s.Stop()
	} else {
		s.StartAsync()
	}

	return err
}
