package elastic

import (
	"errors"
	"fmt"
	"sync"
)

var once sync.Once
var i *instances

type instances struct {
	server []serverInfo
}

type serverInfo struct {
	key      string
	ip 		 string  	
	hostname string		
	port 	 string		
	status   string  	
	zone	 string  	
	timestamp string  	
	name string 		
}
var ErrNotFound = errors.New("not found")

func (i *instances) AddInstance(instance *serverInfo) {
	i.server = append(i.server, *instance)
}

func FindInstance(key string) bool {
	check := false
	Outer:
		for _, instance := range(i.server) {
			if instance.key == key {
				check = true
				break Outer
			}
	}
	return check
}


func GetInstance(key string) (*serverInfo, error) {
	for _ , server := range i.server {
		if server.key == key {
			return &server, nil
		}
	}
	return nil, ErrNotFound
}

func (s *serverInfo) UpdateInstance(status, timestamp string) {
	s.status = status
	s.timestamp = timestamp
}

func CreateInstance(ip, hostname, zone, timestamp, name, status, port string) *serverInfo {
	return &serverInfo{
		key : fmt.Sprintf("%s:%s", ip, port),
		ip : ip,
		name: name,
		hostname:  hostname,
		port : port,
		status : status,
		zone : zone,
		timestamp: timestamp,
	}
}

func InitInstance() *instances {
	once.Do(func() {
		if i == nil {
			i = &instances{}
		}
	})
	return i
}