package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

func GetProperties() {
	m := make(map[interface{}]interface{})
	filename, _ := filepath.Abs("/Users/sn/workspace/goswagger/config/elastic.yaml")
	yamFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamFile, &m)
	if err != nil {
		log.Fatalf("error : %v", err)
	}
	fmt.Println(m)
}