package utils

import (
	"time"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func RFCtoKST(timestr string) string {
	utc, _ := time.Parse(time.RFC3339, timestr)
	loc, _ := time.LoadLocation("Asia/Seoul")
	return utc.In(loc).String()
}