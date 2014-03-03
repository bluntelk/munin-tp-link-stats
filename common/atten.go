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


		"upstream.label Upstream Line Attenuation",
		"upstream.info Line attenuation about 58 is out of spec according to Telstra",
		"upstream.warning 58",
		"upstream.critical 60",

		"downstream.label Downstream Line Attenuation",
		"downstream.info Line attenuation about 58 is out of spec according to Telstra",
		"downstream.warning 58",
		"downstream.critical 60",

	}

	return config
}

func (h *AttenHandlers) GetData() []string {

	modem_data := FetchData(GetModemUrl())

	var upstream_value, downstream_value string

	if upstream_value = fmt.Sprintf("%0.2f", modem_data.Atten.Up); upstream_value == "0.00" {
		upstream_value = "U"
	}
	if downstream_value = fmt.Sprintf("%0.2f", modem_data.Atten.Down); downstream_value == "0.00" {
		downstream_value = "U"
	}

	var data = []string{
		fmt.Sprintf("upstream.value %s", upstream_value),
		fmt.Sprintf("downstream.value %s", downstream_value),
	}
	return data
}
