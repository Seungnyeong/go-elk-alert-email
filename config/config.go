package config

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	"gopkg.in/yaml.v2"
)

// DatabaseConfig Type
type DatabaseConfig struct {
	Name    	string `yaml:"name"`
	User        string `yaml:"user"`
	Password    string `yaml:"password"`
	Host        string `yaml:"host"`
}
 
// ServerConfig Type
type ElasticConfig struct {
	Hosts		[]string 	`yaml:"hosts"`
	Username    string 	 	`yaml:"username"`
	Password  	string    	`yaml:"password"`
	CertPath 	string   	`yaml:"certpath"`
}

// Mail Type

type MailConfig struct {
	Host	string		`yaml:"host"`
	From	string		`yaml:"from"`
}
 
// Config Type
type Config struct {
	Database 	DatabaseConfig   `yaml:"mysql"`
	Elastic   	ElasticConfig    `yaml:"elastic"`
	Mail		MailConfig		 `yaml:"mail"`
}

var once sync.Once
var p *Config

func Properties() *Config {
	once.Do(func() {
		if p == nil {
			p = new(Config)
			filename, _ := filepath.Abs("properties.yaml")
			yamFile, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			err = yaml.Unmarshal(yamFile, &p)
			if err != nil {
				log.Fatalf("error : %v", err)
			}
		}
	})
	return p
}