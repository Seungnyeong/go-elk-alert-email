package config

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"
	"test/utils"

	"gopkg.in/yaml.v2"
)

// DatabaseConfig Type
type DatabaseConfig struct {
	Name     string `yaml:"name"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
}

// ServerConfig Type
type ElasticConfig struct {
	Hosts    []string `yaml:"hosts"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	CertPath string   `yaml:"certpath"`
}

// Mail Type

type MailConfig struct {
	Host string `yaml:"host"`
	From string `yaml:"from"`
	To   string `yaml:"to"`
}

// Config Type
type Config struct {
	Database DatabaseConfig `yaml:"mysql"`
	Elastic  ElasticConfig  `yaml:"elastic"`
	Mail     MailConfig     `yaml:"mail"`
}

var once sync.Once
var P *Config

func Init(path string) *Config {
	once.Do(func() {
		if P == nil {
			P = new(Config)
			filename, _ := filepath.Abs(path)
			yamFile, err := ioutil.ReadFile(filename)
			if err != nil {
				log.Fatal(err)
			}
			err = yaml.Unmarshal(yamFile, &P)
			utils.CheckError(err)
		}
	})
	return P
}
