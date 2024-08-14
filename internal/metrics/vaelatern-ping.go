package metrics

import (
	"time"

	"github.com/mitchellh/mapstructure"
	probing "github.com/prometheus-community/pro-bing"
)

type pingData struct {
	Desc        string `mapstructure:"desc"`
	Host        string `mapstructure:"host"`
	FailIsError bool   `mapstructure:"failIsError"`
}

func pingCard(data interface{}) Card {
	var myOrders pingData
	mapstructure.Decode(data, &myOrders)
	if myOrders.Desc == "" {
		myOrders.Desc = myOrders.Host
	}
	pinger, err := probing.NewPinger(myOrders.Host)
	if err != nil {
		return Card{Title: myOrders.Desc,
			Metric:             err.Error(),
			IsError:            true,
			ShowHeartPulseIcon: true}
	}
	pinger.Count = 1
	pinger.Timeout = time.Second
	pinger.RecordRtts = true
	err = pinger.Run() // Blocks until finished.
	if err != nil {
		return Card{Title: myOrders.Desc,
			Metric:             err.Error(),
			IsError:            true,
			ShowHeartPulseIcon: true}
	}
	stats := pinger.Statistics()
	ok := stats.PacketsRecv == stats.PacketsSent
	if ok {
		return Card{Title: myOrders.Desc,
			Metric:             stats.AvgRtt.Round(time.Millisecond).String(),
			IsOK:               ok,
			ShowHeartPulseIcon: true}
	} else {
		return Card{Title: myOrders.Desc,
			Metric:             "---",
			IsError:            myOrders.FailIsError,
			ShowHeartPulseIcon: true}
	}
}

func init() {
	RegisterCardPlugin("vaelatern-ping", pingCard)
}
