package elastic

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"test/utils"
)

var (
	once sync.Once
	is *instances
	ErrNotFound = errors.New("not found")
)
type instances struct {
	server []*Instance
}


// Instance struct
type Instance struct {
	Key      	string `json:"key"`
	Ip 		 	string `json:"ip"`
	Hostname 	string `json:"hostname"`
	Port 	 	string `json:"port"`
	Status   	string `json:"status"`
	Zone	 	string `json:"zone"`
	Timestamp 	string `json:"timestamp"`
	Name 		string `json:"name"`
	Downcount	int	   `json:"downcount"`
}

func (is *instances) AddInstance(i Instance) {
	is.server = append(is.server, createInstance(i))
}

func FindInstance(key string) bool {
	check := false
	Outer:
		for _, i := range is.server {
			if i.Key == key {
				check = true
				break Outer
			}
		}
	return check
}

func createInstance(i Instance) *Instance {
	newInstance := Instance{
		Key : fmt.Sprintf("%s:%s", i.Ip, i.Port),
		Ip : i.Ip,
		Hostname:  i.Hostname,
		Port : i.Port,
		Status: i.Status,
		Zone : i.Zone,
		Timestamp: utils.RFCtoKST(i.Timestamp),
		Name : i.Name,
		Downcount: 0,
	}
	return &newInstance
}

func (i *Instance) UpdateIntance(status, timestamp string) {
	i.Status = status
	i.Timestamp = timestamp
	if status == "down" {
		i.Downcount++
	} 
}

func (i *Instance) UpdateIntanceDownCount(count int) {
	i.Downcount = count
}

func (is *instances) AllInstance() []*Instance {
	return is.server
}

func (is *instances) GetInstance(key string) (*Instance, error) {
	if len(is.server) > 0 {
		for _, i := range is.server {
			if i.Key == key {
				return i, nil
			}
		}
	}
	return nil, ErrNotFound
}


func GetSingleton() *instances {
	once.Do(func() {
		if is == nil {
			is = new(instances)
		}
	})
	return is
}

func ParsingInstance(response map[string]interface{}) Instance{
	var instance Instance
	for _, hit := range response["hits"].(map[string]interface{})["hits"].([]interface{}) {
		_source := hit.(map[string]interface{})["_source"]
		monitor := _source.(map[string]interface{})["monitor"]
		observer := _source.(map[string]interface{})["observer"]
		instance.Zone = fmt.Sprintf("%s", observer.(map[string]interface{})["geo"].(map[string]interface{})["name"])
		instance.Hostname = fmt.Sprintf("%s",observer.(map[string]interface{})["hostname"])
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
	for _, bucket := range response["aggregations"].(map[string]interface{})["group_by_monitor.id"].(map[string]interface{})["buckets"].([]interface{}){
		motoringIds = append(motoringIds, fmt.Sprintf("%s",bucket.(map[string]interface{})["key"]))
	}
	return motoringIds
}