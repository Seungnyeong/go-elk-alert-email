package cron

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func Start() {
	s := gocron.NewScheduler(time.UTC)

	s.SingletonMode().Every(5).Second().Do(func ()  {
		fmt.Print("승녕잡 1")
	})

	s.StartBlocking()
}