package elastic

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"test/config"
	"test/utils"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type Elastic struct {
	client *elasticsearch.Client
	config *elasticsearch.Config
}

type ElasticAccess interface {
	Search(query *bytes.Buffer, indexName string) (map[string]interface{}, error)
	Status() bool
}

func Client() *Elastic {
	cert, err := utils.GetReadFile(config.P.Elastic.CertPath)
	utils.CheckError(err)
	es := &Elastic{
		config: &elasticsearch.Config{
			Addresses: config.P.Elastic.Hosts,
			Username:  config.P.Elastic.Username,
			Password:  config.P.Elastic.Password,
			CACert:    cert,
		},
	}
	es.client, err = elasticsearch.NewClient(*es.config)
	utils.CheckError(err)
	return es
}

func MakeServerMonitoringQuery(monitorId string) bytes.Buffer {
	body := map[string]interface{}{
		"size": 1,
		"sort": map[string]interface{}{
			"@timestamp": "desc",
		},
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"monitor.id": monitorId,
			},
		},
	}
	return utils.SerializeToJson(body)
}

func MakeServerGroupQuery(ipv4 string) bytes.Buffer {
	filter := []map[string]interface{}{}
	term := map[string]interface{}{
		"term": map[string]interface{}{
			"monitor.ip": ipv4,
		},
	}

	timerange := map[string]interface{}{
		"range": map[string]interface{}{
			"@timestamp": map[string]interface{}{
				"gte":       "now-1d/d",
				"lt":        "now/d",
				"time_zone": "+09:00",
			},
		},
	}

	filter = append(filter, term, timerange)

	body := map[string]interface{}{
		"_source": false,
		"query": map[string]interface{}{
			"bool": map[string]interface{}{
				"filter": filter,
			},
		},
		"aggs": map[string]interface{}{
			"group_by_monitor.id": map[string]interface{}{
				"terms": map[string]interface{}{
					"field": "monitor.id",
				},
			},
		},
	}
	return utils.SerializeToJson(body)
}

func elasticError(res *esapi.Response) error {
	var e map[string]interface{}
	var errMsg string
	if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
		errMsg = err.Error()
	} else {
		errMsg = fmt.Sprintf("[%s] %s: %s", res.Status(), e["error"].(map[string]interface{})["type"], e["error"].(map[string]interface{})["reason"])
	}
	return errors.New(errMsg)
}

func (es Elastic) Status() bool {
	status := true
	res, err := es.client.Info()
	utils.CheckError(err)
	defer res.Body.Close()

	if res.IsError() {
		log.Fatalf("Error: %s", res.String())
		status = false
	}

	return status
}

func (es Elastic) Search(query *bytes.Buffer, indexName string) (map[string]interface{}, error) {
	var r map[string]interface{}
	res, err := es.client.Search(
		es.client.Search.WithContext(context.Background()),
		es.client.Search.WithIndex(indexName),
		es.client.Search.WithBody(query),
		es.client.Search.WithTrackTotalHits(true),
		es.client.Search.WithPretty(),
	)
	defer res.Body.Close()

	if err != nil {
		log.Panicf("Error getting response: %s", err)
	}

	if res.IsError() {
		elasticError(res)
	}

	err = json.NewDecoder(res.Body).Decode(&r)

	if err != nil {
		log.Panicf("Cannot Decode reponse: %s", err)
	}

	return r, err
}
