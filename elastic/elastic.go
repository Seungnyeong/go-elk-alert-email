package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)


type FailedKey struct {
	Username 		string
	Key      		string
	SourceIP		string
	DestinationIP	string
	Type			string
	Category		string
	FailedCount     int
}




func Test() {
	log.SetFlags(0)

  	var (
    	r  map[string]interface{}
    	// wg sync.WaitGroup
  	)
	cert, err := os.ReadFile("cert/ca/ca.crt")
	if err != nil {
		log.Fatal(err)
	}
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
	if err != nil {
		log.Fatalf("Error creating the clinet : %s", err)
	}
	
	res, err := es.Info()
	if err != nil {
    	log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
  
	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
	}

	if err := json.NewDecoder(res.Body).Decode(&r); err != nil {
   	 	log.Fatalf("Error parsing the response body: %s", err)
  	}

	log.Printf("Client: %s", elasticsearch.Version)
  	log.Printf("Server: %s", r["version"].(map[string]interface{})["number"])
  	log.Println(strings.Repeat("~", 37))

	
	var buf bytes.Buffer
	var down string;
	
	for {
		query := map[string]interface{}{
		"size" : 1,
		"sort": map[string]interface{}{
			"@timestamp" : "desc",
		},
		"query" : map[string]interface{}{
			"match" : map[string]interface{}{
				"monitor.id" : "wkmshttp",
			},
		},
		}

		if err := json.NewEncoder(&buf).Encode(query); err != nil {
			log.Fatalf("Error encoding query: %s", err)
		}
		time.Sleep(1 * time.Second)
		res, err := es.Search(
			es.Search.WithContext(context.Background()),
			es.Search.WithIndex("wmp-wkms-health-*"),
			es.Search.WithBody(&buf),
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
			summary := _source.(map[string]interface{})["summary"]
			
			down = fmt.Sprintf("%v",summary.(map[string]interface{})["down"])
		}
		fmt.Println(down)
		res.Body.Close()	
	}
}