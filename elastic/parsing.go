package elastic

import (
	"fmt"
	"strings"
	"test/instances/instance"
)

func ParsingInstance(response map[string]interface{}) instance.Instance {
	var instance instance.Instance
	for _, hit := range response["hits"].(map[string]interface{})["hits"].([]interface{}) {
		_source := hit.(map[string]interface{})["_source"]
		monitor := _source.(map[string]interface{})["monitor"]
		observer := _source.(map[string]interface{})["observer"]
		instance.Zone = fmt.Sprintf("%s", observer.(map[string]interface{})["geo"].(map[string]interface{})["name"])
		instance.Hostname = fmt.Sprintf("%s", observer.(map[string]interface{})["hostname"])
		instance.Timestamp = fmt.Sprintf("%s", _source.(map[string]interface{})["@timestamp"])
		instance.Ip = fmt.Sprintf("%s", monitor.(map[string]interface{})["ip"])
		instance.Status = fmt.Sprintf("%s", monitor.(map[string]interface{})["status"])
		instance.Port = strings.Split(fmt.Sprintf("%f", _source.(map[string]interface{})["url"].(map[string]interface{})["port"]), ".")[0]
		instance.Name = fmt.Sprintf("%s", monitor.(map[string]interface{})["name"])
	}
	return instance
}

func ParsingInstanceId(response map[string]interface{}) []string {
	var motoringIds []string
	for _, bucket := range response["aggregations"].(map[string]interface{})["group_by_monitor.id"].(map[string]interface{})["buckets"].([]interface{}) {
		motoringIds = append(motoringIds, fmt.Sprintf("%s", bucket.(map[string]interface{})["key"]))
	}
	return motoringIds
}
