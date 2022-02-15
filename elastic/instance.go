package elastic

import (
	"errors"
	"fmt"
	"sync"
)

var (
	once sync.Once
	is *instances
	ErrNotFound = errors.New("not found")
)
type instances struct {
	server []*Instance
}

type Instance struct {
	Key      	string `json:"key"`
	Ip 		 	string `json:"Ip"`
	Hostname 	string `json:"hostname"`
	Port 	 	string `json:"port"`
	Status   	string `json:"status"`
	Zone	 	string `json:"zone"`
	Timestamp 	string `json:"timestamp"`
	Name 		string `json:"name"`
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
		Timestamp: i.Timestamp,
		Name : i.Name,
	}
	return &newInstance
}

func (i *Instance) UpdateIntance(status, timestamp string) {
	i.Status = status
	i.Timestamp = timestamp
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