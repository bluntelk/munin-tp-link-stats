package main

import (
	"fmt"
	"github.com/bluntelk/munin-tp-link-stats/common"
)

type handlers struct {}

func (h *handlers) GetAutoConfig() bool {
	return false
}

func (h *handlers) GetConfig() []string {
	var config = []string{
		"graph_args --base 1000 -l 0 --upper-limit 90",
		"graph_vlabel SNR db",
		"graph_title Signal to Noise Ratio",
		"snr.label Signal to Noise Ratio",
		"snr.info SNR below 6 is bad it means that there is too much noise for a good connection. You want a steady straight line with little variation.",
	}

	return config
}

func (h *handlers) GetData() []string {

	modem_data := common.FetchData(common.GetModemUrl())

	var data = []string{
		fmt.Sprintf("snr.up %0.2f", modem_data.Snr.Up),
		fmt.Sprintf("snr.down %0.2f", modem_data.Snr.Down),
	}
	return data
}

func main() {

	h := new(handlers)
	action := common.GetAction()
	response := common.HandleAction(action, h)
	common.ShowResponse(response)
}
