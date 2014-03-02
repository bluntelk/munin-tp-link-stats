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
		"sync.label Line Sync Data Rate kbps",
		"sync.info The current sync speed to the Internet",

	}

	return config
}

func (h *SyncHandlers) GetData() []string {

	modem_data := FetchData(GetModemUrl())

	var data = []string{
		fmt.Sprintf("sync.up %0.0f", modem_data.Sync.Up),
		fmt.Sprintf("sync.down %0.0f", modem_data.Sync.Down),
		fmt.Sprintf("sync.max_up %0.0f", modem_data.MaxRate.Up),
		fmt.Sprintf("sync.max_down %0.0f", modem_data.MaxRate.Down),
	}
	return data
}
