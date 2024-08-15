package metrics

import (
	"net/http"
	"slices"
	"strconv"

	"github.com/mitchellh/mapstructure"
)

type vaelatern_http200Data struct {
	Desc     string `mapstructure:"desc"`
	Url      string `mapstructure:"url"`
	Verb     string `mapstructure:"http-verb"`
	StatusOK []int  `mapstructure:"status-codes-accepted"`
}

func vaelatern_http200Card(data interface{}) Card {
	var myOrders vaelatern_http200Data
	mapstructure.Decode(data, &myOrders)
	if myOrders.Verb == "" {
		myOrders.Verb = "GET"
	}
	if len(myOrders.StatusOK) == 0 {
		myOrders.StatusOK = []int{200}
	}
	if myOrders.Url == "" {
		return Card{Title: myOrders.Desc,
			IconText: "HTTP",
			Metric:   "No URL To Test",
			IsError:  true}
	}

	req, err := http.NewRequest(myOrders.Verb, myOrders.Url, nil)
	if err != nil {
		return Card{Title: myOrders.Desc,
			IconText: "HTTP",
			Metric:   err.Error(),
			IsError:  true}
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return Card{Title: myOrders.Desc,
			IconText: "HTTP",
			Metric:   err.Error(),
			IsError:  true}
	}

	ok := slices.Contains(myOrders.StatusOK, resp.StatusCode)
	if ok {
		return Card{Title: myOrders.Desc,
			IconText: "HTTP",
			IsOK:     true}
	} else {
		return Card{Title: myOrders.Desc,
			IconText: "HTTP",
			Metric:   strconv.Itoa(resp.StatusCode),
			IsError:  true}
	}
}

func init() {
	RegisterCardPlugin("vaelatern-http-200", vaelatern_http200Card)
}
