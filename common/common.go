package common

import (
	"fmt"
	"os"
	"net/http"
	"io/ioutil"
	"strings"
	"regexp"
	"strconv"
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

func (m *ModemStats) Populate(page string) {
	skitch_stat(page, "(?s)SNR Margin.*?</tr>", "^(.*) ([0-9\\.]+) ([0-9\\.]+) (db)", &m.Snr)
	skitch_stat(page, "(?s)Line Attenuation.*?</tr>", "^(.*) ([0-9\\.]+) ([0-9\\.]+) (db)", &m.Atten)
	skitch_stat(page, "(?s)Data Rate.*?</tr>", "^(.*) ([0-9\\.]+) ([0-9\\.]+) (kbps)", &m.Sync)
	skitch_stat(page, "(?s)Max Rate.*?</tr>", "^(.*) ([0-9\\.]+) ([0-9\\.]+) (kbps)", &m.MaxRate)
	skitch_stat(page, "(?s)POWER.*?</tr>", "^(.*) ([0-9\\.]+) ([0-9\\.]+) (dbm)", &m.Power)
	skitch_stat(page, "(?s)CRC.*?</tr>", "^(.*) ([0-9\\.]+) ([0-9\\.]+)(.*)", &m.Crc)
}

func skitch_stat(page, find_re, grab_re string, stat *UpDownStats) {
	re_stripper := regexp.MustCompile("(?s)<.*?>")
	re_trimmer := regexp.MustCompile("(?s)[\\s:]+")

	re_find := regexp.MustCompile(find_re)
	re_grab := regexp.MustCompile(grab_re)

	found_text := re_find.FindString(page)
	found_text_detagged := re_stripper.ReplaceAllString(found_text, " ")
	found_text_clean := re_trimmer.ReplaceAllString(found_text_detagged, " ")
	fields := re_grab.FindAllStringSubmatch(found_text_clean, 4)

	if len(fields) > 0 {
		//var error error
		stat.Down, _ = strconv.ParseFloat(fields[0][2],64)
		stat.Up, _ = strconv.ParseFloat(fields[0][3],64)
		stat.Unit = fields[0][4]
		PluginDebugPrint(fmt.Sprintf("Found Stat: %+v", stat))
	}
}

type ModemUrl struct {
	Host, Path, Username, Password string
}

func (m *ModemUrl) AsUrl() string {
	return fmt.Sprintf("http://%s:%s@%s%s", m.Username, m.Password, m.Host, m.Path)
}

func GetModemUrl() ModemUrl {
	var m ModemUrl
	m.Host = GetValueFromEnv("tplink_host", "")
	m.Path = GetValueFromEnv("tplink_path", "/status/status_deviceinfo.htm")
	m.Username = GetValueFromEnv("tplink_username", "admin")
	m.Password = GetValueFromEnv("tplink_password", "")

	PluginDebugPrint(fmt.Sprintf("Using URL %s", m.AsUrl()))
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
	var data []byte
	client := new(http.Client)
	PluginDebugPrint(fmt.Sprintf("GETting page: %s", url.AsUrl()))
	page, err := client.Get(url.AsUrl())
	if err != nil {
		PluginDebugPrint(fmt.Sprintf("Failed to get modem page, %s", err))
		return stats;
	}
	PluginDebugPrint("Fetched the page")
	data, err = ioutil.ReadAll(page.Body)
	page.Body.Close()
	stats.Populate(string(data))

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
