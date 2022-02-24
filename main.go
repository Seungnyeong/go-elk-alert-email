package main

import (
	"os"
	"test/cli"
	"test/utils"
)

func main() {
	if _ , err := os.Stat("logs"); err != nil {
		merr := os.MkdirAll("logs", os.ModePerm)
		utils.CheckError(merr)
	}
	cli.Start()
}
