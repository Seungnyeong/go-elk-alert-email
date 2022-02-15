package elastic

import (
	"fmt"
	"log"
	"test/utils"
	"time"

	"github.com/go-co-op/gocron"
)

func Job(agentId []string) {
	GetSingleton()
	es, _ := ElasticConnection()
	
	res, err := es.Info()
	utils.CheckError(err)
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	
	for _, name := range agentId {
		result := elsticResult(es, name)
		if !FindInstance(fmt.Sprintf("%s:%s", result.Ip, result.Port)) {
			is.AddInstance(result)		
		} else {
			a, err := is.GetInstance(fmt.Sprintf("%s:%s", result.Ip, result.Port))
			utils.CheckError(err)
			a.UpdateIntance(result.Status, result.Timestamp)
		}
	}
	fmt.Println(is.AllInstance()[0])
	
}

func CronJob() {
	var a = []string{"wkms", "wkmsdb", "wkmshttp"}
	s := gocron.NewScheduler(time.UTC)
	s.SingletonMode().Every(5).Second().Do(func ()  {
		Job(a)
	})
	s.StartAsync()
}