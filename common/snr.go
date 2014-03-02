package common

import (
	"fmt"
)

type SnrHandlers struct {}

func (h *SnrHandlers) GetAutoConfig() bool {
	return false
}

func (h *SnrHandlers) GetConfig() []string {
	var config = []string{
		"graph_args --base 1000 -l 0 --upper-limit 90",
		"graph_vlabel SNR db",
		"graph_scale no",
		"graph_title Signal to Noise Ratio",
		"snr.label Signal to Noise Ratio",
		"snr.info SNR below 6 is bad it means that there is too much noise for a good connection. You want a steady straight line with little variation.",

	}

	return config
}

func (h *SnrHandlers) GetData() []string {

	modem_data := FetchData(GetModemUrl())

	var data = []string{
		fmt.Sprintf("snr.up %0.2f", modem_data.Snr.Up),
		fmt.Sprintf("snr.down %0.2f", modem_data.Snr.Down),
	}
	return data
}
