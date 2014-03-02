package main

import (
	"errors"
	"fmt"
	"github.com/bluntelk/munin-tp-link-stats/common"
	"os"
	"path/filepath"
)

func getValidLookup() map[string]common.ResponseStruct {
	return map[string]common.ResponseStruct{
		"tplink_snr":   new(common.SnrHandlers),
		"tplink_atten": new(common.AttenHandlers),
		"tplink_sync": new(common.SyncHandlers),
		"tplink_power": new(common.PowerHandlers),
		"tplink_crc": new(common.CrcHandlers),
	}
}

func determineHandlers() (common.ResponseStruct, error) {

	lookup := getValidLookup()

	key_file := filepath.Base(os.Args[0])
	key_env := os.Getenv("STAT")

	if response_file, ok_file := lookup[key_file]; ok_file {
		return response_file, nil
	} else if response_env, ok_env := lookup[key_env]; ok_env {
		return response_env, nil
	} else {
		return new(common.SnrHandlers), errors.New(fmt.Sprintf("filename:%s or env%%STAT='%s' are not a valid stat type.", key_file, key_env))
	}
}

func showUsage() {
	fmt.Println("This program is designed to be symlinked to one of the following")
	for key, _ := range getValidLookup() {
		fmt.Printf("\t* %s\n", key)
	}
	fmt.Println("The following environment variables are used: STAT, tplink_username, tplink_password, tplink_host, tplink_path")
}

func main() {

	h, err := determineHandlers()
	if err != nil {
		fmt.Println(err)
		showUsage()
	} else {
		action := common.GetAction()
		response := common.HandleAction(action, h)
		common.ShowResponse(response)
	}
}
