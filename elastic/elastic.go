package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"test/utils"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)



func ElasticConnection() (*elasticsearch.Client, error ){
	cert, err := os.ReadFile("cert/ca/ca.crt")
	utils.CheckError(err)
	cfg := elasticsearch.Config{
		Addresses: []string{
			"https://172.19.31.75:9200",
			"https://172.19.31.98:9200",
			"https://172.19.31.79:9200",
			"https://172.19.31.60:9200",
			"https://172.19.31.101:9200",
		},
		Username: "elastic",
		Password: "EuVt9rRDLEnv2f6XwHTW",
		CACert: cert,
	}

	es, err := elasticsearch.NewClient(cfg)
	utils.CheckError(err)

	return es, err
}

func elasticQuery(monitorId string) bytes.Buffer {
	var query bytes.Buffer
	body := map[string]interface{}{
		"size" : 1,
		"sort": map[string]interface{}{
			"@timestamp" : "desc",
		},
		"query" : map[string]interface{}{
			"match" : map[string]interface{}{
				"monitor.id" : monitorId,
			},
		},
	}
	err := json.NewEncoder(&query).Encode(body)
	utils.CheckError(err)
	return query
}

func getInstanceStatus(es *elasticsearch.Client, r map[string]interface{}, agentId string) string {
	var down string
	query := elasticQuery(agentId)
		time.Sleep(5 * time.Second)
		res, err := es.Search(
			es.Search.WithContext(context.Background()),
			es.Search.WithIndex("wmp-wkms-health-*"),
			es.Search.WithBody(&query),
			es.Search.WithTrackTotalHits(true),
			es.Search.WithPretty(),
		)

		if err != nil {
			log.Fatalf("Error getting response: %s", err)
		}
		
		if res.IsError() {
			var e map[string]interface{}
			if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
				log.Fatalf("Error parsing the response body: %s", err)
			} else {
				log.Fatalf("[%s] %s: %s",
					res.Status(),
					e["error"].(map[string]interface{})["type"],
					e["error"].(map[string]interface{})["reason"],
				)
			}
		}

		if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		}
		var instance *instance
		for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
			_source := hit.(map[string]interface{})["_source"]
			summary := _source.(map[string]interface{})["summary"]
			monitor := _source.(map[string]interface{})["monitor"]
			observer := _source.(map[string]interface{})["observer"]
			zone := fmt.Sprintf("%s", observer.(map[string]interface{})["geo"].(map[string]interface{})["name"])
			hostname := fmt.Sprintf("%s",observer.(map[string]interface{})["hostname"])
			timestamp := fmt.Sprintf("%s", _source.(map[string]interface{})["@timestamp"])
			ip := fmt.Sprintf("%s", monitor.(map[string]interface{})["ip"])
			status := fmt.Sprintf("%s", monitor.(map[string]interface{})["status"])
			port, _ := strconv.Atoi(fmt.Sprintf("%f", _source.(map[string]interface{})["url"].(map[string]interface{})["port"]))
			down = fmt.Sprintf("%v",summary.(map[string]interface{})["down"])
			name := fmt.Sprintf("%s", monitor.(map[string]interface{})["name"])
			fmt.Println(ip, status, zone, port, timestamp, hostname, name)
			if !FindInstance(name) {
				instance = CreateInstance(ip, hostname, zone, timestamp, name, status, port)
			}
			
			if instance != nil {
				i.AddInstance(instance)
			}

		}	
		defer res.Body.Close()
		fmt.Println(i)
		return down
}

func Start() {
	log.SetFlags(0)
	i = InitInstance()
	var (
		r map[string]interface{}
	)

	es, _ := ElasticConnection()
	res, err := es.Info()
	utils.CheckError(err)
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	
	for {
		nginx := getInstanceStatus(es, r, "wkmshttp")
		uwsgi := getInstanceStatus(es, r, "wkms")
		mysql := getInstanceStatus(es, r, "wkmsdb")
		fmt.Printf("[Status Nginx] : %s\n",nginx)
		fmt.Printf("[Status uwsgi] : %s\n",uwsgi)
		fmt.Printf("[Status mysql] : %s\n",mysql)
	}
}
