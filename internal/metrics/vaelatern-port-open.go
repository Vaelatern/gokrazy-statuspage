package metrics

import (
	"net"
	"strconv"
	"time"

	"github.com/mitchellh/mapstructure"
)

type vaelatern_portData struct {
	Desc  string `mapstructure:"desc"`
	Host  string `mapstructure:"host"`
	Port  int    `mapstructure:"port"`
	Proto string `mapstructure:"proto"`
}

func vaelatern_portCard(data interface{}) Card {
	var myOrders vaelatern_portData
	mapstructure.Decode(data, &myOrders)
	if myOrders.Desc == "" {
		myOrders.Desc = myOrders.Host
	}
	if myOrders.Proto == "" {
		myOrders.Proto = "tcp"
	}
	conn, err := net.DialTimeout(myOrders.Proto, net.JoinHostPort(myOrders.Host, strconv.Itoa(myOrders.Port)), time.Second)
	if err != nil {
		return Card{Title: myOrders.Desc,
			Metric:         err.Error(),
			IsError:        true,
			ShowServerIcon: true}
	}
	if conn != nil {
		return Card{Title: myOrders.Desc,
			IsOK:           true,
			ShowServerIcon: true}
		defer conn.Close()
	}
	return Card{Title: myOrders.Desc,
		Metric:         "???",
		ShowServerIcon: true}
}

func init() {
	RegisterCardPlugin("vaelatern-port-open", vaelatern_portCard)
}
