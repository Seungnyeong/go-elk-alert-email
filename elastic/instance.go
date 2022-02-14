package elastic

import (
	"sync"
)

var once sync.Once
var i *instances

type instances struct {
	server []instance
}

type instance struct {
	Ip 		 string  	`json"ip"`
	Hostname string		`json"hostname"`
	Port 	 int		`json"port"`
	Status   string  	`json"status"`
	Zone	 string  	`json"zone"`
	Timestamp string  	`json"timestamp"`
	Name string 		`json"name"`
}


func (i *instances) AddInstance( instance *instance) {
	i.server = append(i.server, *instance)
}

func FindInstance(name string) bool {
	check := false
	Outer:
		for _, instance := range(i.server) {
			if instance.Name == name {
				check = true
				break Outer
			}
	}
	return check
}

func CreateInstance(ip, hostname, zone, timestamp, name, status string, port int) *instance {
	return &instance{
		Ip : ip,
		Name: name,
		Hostname:  hostname,
		Port : port,
		Status : status,
		Zone : zone,
		Timestamp: timestamp,
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