package instance

import (
	"test/utils"
)

// Instance struct
type Instance struct {
	Ip        string `json:"ip"`
	Hostname  string `json:"hostname"`
	Port      string `json:"port"`
	Status    string `json:"status"`
	Zone      string `json:"zone"`
	Timestamp string `json:"timestamp"`
	Name      string `json:"name"`
	Downcount int    `json:"downcount"`
	Mailed    bool   `json:"mailed"`
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

func (i *Instance) UpdateMailed() {
	i.Mailed = !i.Mailed
}

func CreateInstance(i Instance) *Instance {
	newInstance := Instance{
		Ip:        i.Ip,
		Hostname:  i.Hostname,
		Port:      i.Port,
		Status:    i.Status,
		Zone:      i.Zone,
		Timestamp: utils.RFCtoKST(i.Timestamp),
		Name:      i.Name,
		Downcount: 0,
		Mailed:    false,
	}
	return &newInstance
}
