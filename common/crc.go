package common

import (
	"fmt"
)

type CrcHandlers struct {}

func (h *CrcHandlers) GetAutoConfig() bool {
	return false
}

func (h *CrcHandlers) GetConfig() []string {
	var config = []string{
		"graph_args --base 1000 -l 0",
		"graph_vlabel Line CRC Errors",
		"graph_scale no",
		"graph_title CRC Errors",

		"upstream.label Upstream CRC Errors",
		"upstream.info The number of Upstream CRC errors",

		"downstream.label Downstream CRC Errors",
		"downstream.info The number of Downstream CRC errors",

	}

	return config
}

func (h *CrcHandlers) GetData() []string {

	modem_data := FetchData(GetModemUrl())

	var data = []string{
		fmt.Sprintf("upstream.value %0.2f", modem_data.Crc.Up),
		fmt.Sprintf("downstream.value %0.2f", modem_data.Crc.Down),
	}
	return data
}
