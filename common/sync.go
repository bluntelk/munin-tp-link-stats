package common

import (
	"fmt"
)

type SyncHandlers struct {}

func (h *SyncHandlers) GetAutoConfig() bool {
	return false
}

func (h *SyncHandlers) GetConfig() []string {
	var config = []string{
		"graph_args --base 1000 -l 0 --upper-limit 25000",
		"graph_vlabel Line Sync Data Rate kbps",
		"graph_scale no",
		"graph_title Line Sync Data Rate kbps",

		"upstream.label Upstream Line Sync Data Rate kbps",
		"upstream.info The current upstream sync speed to the Internet",

		"downstream.label Downstream Line Sync Data Rate kbps",
		"downstream.info The current downstream sync speed to the Internet",

		"max_upstream.label Max Upstream Line Sync Data Rate kbps",
		"max_upstream.info Max current upstream sync speed to the Internet",

		"max_downstream.label Max Downstream Line Sync Data Rate kbps",
		"max_downstream.info Max current downstream sync speed to the Internet",

	}

	return config
}

func (h *SyncHandlers) GetData() []string {

	modem_data := FetchData(GetModemUrl())

	var data = []string{
		fmt.Sprintf("upstream.value %0.0f", modem_data.Sync.Up),
		fmt.Sprintf("downstream.value %0.0f", modem_data.Sync.Down),
		fmt.Sprintf("max_upstream.value %0.0f", modem_data.MaxRate.Up),
		fmt.Sprintf("max_downstream.value %0.0f", modem_data.MaxRate.Down),
	}
	return data
}
