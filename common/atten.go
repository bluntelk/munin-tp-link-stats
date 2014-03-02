package common

import (
	"fmt"
)

type AttenHandlers struct {}

func (h *AttenHandlers) GetAutoConfig() bool {
	return false
}

func (h *AttenHandlers) GetConfig() []string {
	var config = []string{
		"graph_args --base 1000 -l 0 --upper-limit 90",
		"graph_vlabel Line Attenuation db",
		"graph_scale no",
		"graph_title Line Attenuation",
		"atten.label Line Attenuation",
		"atten.info Line attenuation about 58 is out of spec according to Telstra",
		"atten.warning 58",
		"atten.critical 60",

	}

	return config
}

func (h *AttenHandlers) GetData() []string {

	modem_data := FetchData(GetModemUrl())

	var data = []string{
		fmt.Sprintf("atten.up %0.2f", modem_data.Atten.Up),
		fmt.Sprintf("atten.down %0.2f", modem_data.Atten.Down),
	}
	return data
}
