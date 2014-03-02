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

		"upstream.label Upstream Power dbm",
		"upstream.info The amount of power your upstream connection is using",

		"downstream.label Downstream Power dbm",
		"downstream.info The amount of power your downstream connection is using",

	}

	return config
}

func (h *PowerHandlers) GetData() []string {

	modem_data := FetchData(GetModemUrl())

	var data = []string{
		fmt.Sprintf("upstream.value %0.2f", modem_data.Power.Up),
		fmt.Sprintf("downstream.value %0.2f", modem_data.Power.Down),
	}
	return data
}
