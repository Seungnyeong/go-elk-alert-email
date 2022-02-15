package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"test/utils"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)
var eo = new(ElasticOutPut)
var a = [3]string{"wkms", "wkmsdb", "wkmshttp"}
type ElasticOutPut struct {
	zone  string
	hostname string
	timestamp string
	ip		string
	status	string
	port	string
	name	string
}


func ElasticConnection() (*elasticsearch.Client, error ){
	cert, err := ioutil.ReadFile("cert/ca/ca.crt")
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

func getInstanceStatus(es *elasticsearch.Client, agentId string) *ElasticOutPut {
	var (
		r map[string]interface{}
	)

	query := elasticQuery(agentId)
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
		
		for _, hit := range r["hits"].(map[string]interface{})["hits"].([]interface{}) {
			_source := hit.(map[string]interface{})["_source"]
			monitor := _source.(map[string]interface{})["monitor"]
			observer := _source.(map[string]interface{})["observer"]
			eo.zone = fmt.Sprintf("%s", observer.(map[string]interface{})["geo"].(map[string]interface{})["name"])
			eo.hostname = fmt.Sprintf("%s",observer.(map[string]interface{})["hostname"])
			eo.timestamp = fmt.Sprintf("%s", _source.(map[string]interface{})["@timestamp"])
			eo.ip = fmt.Sprintf("%s", monitor.(map[string]interface{})["ip"])
			eo.status = fmt.Sprintf("%s", monitor.(map[string]interface{})["status"])
			eo.port = strings.Split(fmt.Sprintf("%f", _source.(map[string]interface{})["url"].(map[string]interface{})["port"]), ".")[0]
			eo.name = fmt.Sprintf("%s", monitor.(map[string]interface{})["name"])
		}	

		defer res.Body.Close()
		return eo
}

func (i *instances) Start() {
	log.SetFlags(0)
	es, _ := ElasticConnection()
	res, err := es.Info()
	utils.CheckError(err)
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}
	
	var instance *serverInfo
	for {
			time.Sleep(3 * time.Second)	
			t := getInstanceStatus(es, a[0])
			if !FindInstance(fmt.Sprintf("%s:%s", t.ip, t.port)) {
				instance = CreateInstance(t.ip, t.hostname, t.zone, t.timestamp, t.name, t.status, t.port)
				i.AddInstance(instance)
			}
			update , err := GetInstance(fmt.Sprintf("%s:%s", t.ip, t.port))
			utils.CheckError(err)
			update.UpdateInstance(t.status, t.timestamp)
			fmt.Println(i)
		}
	
}
