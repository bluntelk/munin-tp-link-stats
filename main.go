package main

import (
	"fmt"
	"os"
	"strings"
)

func handleAction(action string) string {
	ret := ""
	if action == "autoconf" {
		ret = "yes"
	} else if action == "config" {
		var config = []string{
			"multigraph snr",

			"multigraph atten",
			"graph_args --base 1000 -l 0 --upper-limit 25000",
			"graph_vlabel mbits",
			"graph_title Internet Connection Sync Speed",


			"multigraph sync",
			"graph_args --base 1000 -l 0 --upper-limit 25000",
			"graph_vlabel mbits",
			"graph_title Internet Connection Sync Speed",

			"multigraph sync_max",
			"graph_args --base 1000 -l 0 --upper-limit 25000",
			"graph_vlabel mbits",
			"graph_title Max Sync speed",

			"multigraph power",
			"graph_args --base 1000 -l 0 --upper-limit 50",
			"graph_vlabel dbm",
			"graph_title Power",

			"multigraph crc",
			"graph_args --base 1000 -l 0",
			"graph_vlabel dbm",
			"graph_title Power",




			"graph_args --base 1000 -l 0 --upper-limit 25000",
			"graph_vlabel mbits",
			"graph_title Internet Connection",
			"graph_category network",
			"graph_info Shows the historical connection health of the modem",
			"graph_order snr, atten, sync, sync_max, power, crc, status",

			"snr.label Signal to Noise Ratio",
			"snr.info SNR below 6 is bad it means that there is too much noise for a good connection. You want a steady straight line with little variation.",

			"atten.label Signal Attenuation",

			"sync.label Data Sync Speed",

			"sync_max.label Max Sync Speed",

			"power.label Power",

			"crc.label CRC",

			"status.label Whether or not we are online",
		}
		ret = strings.Join(config, "\n")
	}
	return ret
}

func main() {
	fmt.Println(handleAction(os.Args[1]))
}
