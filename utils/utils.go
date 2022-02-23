package utils

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"time"
)

func CheckError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func RFCtoKST(timestr string) string {
	utc, _ := time.Parse(time.RFC3339, timestr)
	loc, _ := time.LoadLocation("Asia/Seoul")
	return utc.In(loc).String()
}

func GetReadFile(dir string) ([]byte, error) {
	cert, err := ioutil.ReadFile(dir)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

func SerializeToJson(body interface{}) bytes.Buffer {
	var query bytes.Buffer
	err := json.NewEncoder(&query).Encode(body)
	CheckError(err)
	return query
}

func CheckIPAddress(ip string) bool {
	if net.ParseIP(ip) == nil {
		return false
	} else {
		return true
	}
}
