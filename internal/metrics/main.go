package metrics

import (
	"sync"

	"github.com/mitchellh/mapstructure"
)

type Card struct {
	Title              string
	Metric             string
	IconText           string
	ShowHeartPulseIcon bool
	ShowServerIcon     bool
	IsOK               bool
	IsError            bool
}

type Callback func(interface{}) Card

var (
	plugins map[string]Callback
)

func init() {
	plugins = make(map[string]Callback)
}

func RegisterCardPlugin(name string, cb Callback) {
	plugins[name] = cb
}

type inputData struct {
	Plugin string `mapstructure:"type"`
}

func AllCards(data []interface{}) []Card {
	var returnVal []Card = make([]Card, len(data))
	var wg sync.WaitGroup
	cardChan := make(chan struct {
		int
		Card
	}, len(data))
	for i, d := range data {
		wg.Add(1)
		go func(i int, d interface{}) {
			defer wg.Done()
			var input inputData
			mapstructure.Decode(d, &input)
			if whichPlugin, found := plugins[input.Plugin]; found {
				cardChan <- struct {
					int
					Card
				}{i, whichPlugin(d)}
			} else {
				cardChan <- struct {
					int
					Card
				}{i, Card{Title: "Unknown type"}}
			}
		}(i, d)
	}
	wg.Wait()
	close(cardChan)
	for v := range cardChan {
		returnVal[v.int] = v.Card
	}
	return returnVal
}
