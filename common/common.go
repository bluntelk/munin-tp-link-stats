package common

import (
	"fmt"
	"os"
	"net/http"
	"ioutil"
	"strings"
)

type ResponseStruct interface {
	GetAutoConfig() bool
	GetConfig() []string
	GetData() []string
}

type UpDownStats struct {
	Up, Down, Min, Max float64
	Unit               string
}
type ModemStats struct {
	Snr     UpDownStats
	Atten   UpDownStats
	Sync    UpDownStats
	MaxRate UpDownStats
	Power   UpDownStats
	Crc     UpDownStats
}
func (m *ModemStats)Populate(page string) {

}
type ModemUrl struct {
	Host, Path, Username, Password string
}

func (m *ModemUrl)AsUrl() string {
	return fmt.Sprintf("http://%s:%s@%s%s", m.Username, m.Password, m.Host, m.Path)
}

func GetModemUrl() ModemUrl {
	var m ModemUrl
	m.Host = GetValueFromEnv("tplink_host", "")
	m.Path = GetValueFromEnv("tplink_path", "/status/status_deviceinfo.htm")
	m.Username = GetValueFromEnv("tplink_username", "admin")
	m.Password = GetValueFromEnv("tplink_password", "")

	PluginDebugPrint(fmt.Sprintf("Using URL %s",m.AsUrl()))
	return m
}

func GetValueFromEnv(value, default_value string) string {
	ret_value := os.Getenv(value)
	if ret_value == "" {
		ret_value = default_value
	}
	return ret_value
}

func ShowResponse(response string) {
	fmt.Println(response)
}

func GetAction() string {
	action := "data"
	if len(os.Args) > 1 {
		action = os.Args[1]
	}

	PluginDebugPrint(fmt.Sprintf("Command Line Action: %s", action))
	return action
}

func FetchData(url ModemUrl) ModemStats {
	var stats ModemStats
	client := new(http.Client)
	data, err := client.Get(url.AsUrl())
	if err != nil {
		PluginDebugPrint(fmt.Sprintf("Failed to get modem page, %s", err))
		return stats;
	}
	stats.Populate(data)

	return stats;
}

func HandleAction(action string, handlers ResponseStruct) string {
	ret := ""
	if action == "autoconf" {
		PluginDebugPrint("Performing Action: Auto Config")
		if handlers.GetAutoConfig() {
			ret = "yes"
		} else {
			ret = "no"
		}

	} else if action == "config" {
		PluginDebugPrint("Performing Action: Config")
		ret = strings.Join(handlers.GetConfig(), "\n");

	} else {
		PluginDebugPrint("Performing Action: Gather Data")
		ret = strings.Join(handlers.GetData(), "\n");
	}

	return ret
}

func PluginDebugPrint(message string) {
	if GetValueFromEnv("MUNIN_DEBUG", "") == "1" {
		fmt.Println(fmt.Sprintf("#DEBUG: %s", message))
	}

}
