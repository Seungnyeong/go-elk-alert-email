package main

import (
	"os"
	"test/cli"
	"test/utils"
)

func main() {
	path := utils.GetBinPath()
	if _ , err := os.Stat(path + "/logs"); err != nil {
		merr := os.MkdirAll("logs", os.ModePerm)
		utils.CheckError(merr)
	}
	cli.Start()
}
