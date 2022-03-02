package service

import (
	"errors"
	"fmt"
	"sync"
	"test/instances/instance"
)

var (
	once        sync.Once
	is          *instances
	ErrNotFound = errors.New("not found")
)

type instances struct {
	server map[string]*instance.Instance
	m      sync.Mutex
}

func NewInstances() *instances {
	once.Do(func() {
		if is == nil {
			is = &instances{
				server: make(map[string]*instance.Instance),
			}
		}
	})
	return is
}

func (is *instances) AddInstance(i instance.Instance) *instance.Instance {
	key := fmt.Sprintf("%s:%s", i.Ip, i.Port)
	is.server[key] = instance.CreateInstance(i)
	return is.server[key]
}

func (is *instances) RemoveInstance(key string) {
	is.m.Lock()
	defer is.m.Unlock()
	delete(is.server, key)
}

func (is *instances) GetAllInstance() map[string]*instance.Instance {
	return is.server
}

func (is *instances) GetInstance(key string) (*instance.Instance, error) {
	server, ok := is.server[key]
	if !ok {
		return nil, ErrNotFound
	}
	return server, nil
}

func (is *instances) UpdateDownCount(key string) bool {
	down := false

	if is.server[key].Downcount%10 == 0 && is.server[key].Status == "down" {
		down = true
		is.server[key].UpdateIntanceDownCount(0)
	}

	if is.server[key].Downcount > 0 && is.server[key].Status == "up" {
		is.server[key].UpdateIntanceDownCount(0)
		is.server[key].Mailed = false
	}

	return down
}
