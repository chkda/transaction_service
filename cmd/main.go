package main

import (
	"os"

	service "github.com/chkda/transaction_service/internal"
)

var FILE_LOC = "/config/config.json"

func main() {
	currDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configFile := currDir + FILE_LOC
	service.Start(configFile)
}
