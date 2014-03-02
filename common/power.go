package common

import (
	"fmt"
)

type PowerHandlers struct {}

func (h *PowerHandlers) GetAutoConfig() bool {
	return false
}

func (h *PowerHandlers) GetConfig() []string {
	var config = []string{
		"graph_args --base 1000 -l 0",
		"graph_vlabel Line Power dbm",
		"graph_scale no",
		"graph_title Power dbm",
		"power.label Power dbm",
		"power.info The amount of power in the connection",

	}

	return config
}

func (h *PowerHandlers) GetData() []string {

	modem_data := FetchData(GetModemUrl())

	var data = []string{
		fmt.Sprintf("power.up %0.2f", modem_data.Power.Up),
		fmt.Sprintf("power.down %0.2f", modem_data.Power.Down),
	}
	return data
}
