package metrics

type Card struct {
	Title              string
	Metric             string
	IconText           string
	ShowHeartPulseIcon bool
	ShowServerIcon     bool
	IsOK               bool
	IsError            bool
}

func AllCards() []Card {
	return []Card{
		Card{Title: "Self", Metric: "0ms", IsOK: true, ShowHeartPulseIcon: true},
		Card{Title: "Self", ShowServerIcon: true, IsOK: true},
		Card{Title: "Office1", Metric: "317ms", IsOK: true, ShowHeartPulseIcon: true},
		Card{Title: "Server", Metric: "68ms", IsOK: true, ShowHeartPulseIcon: true},
		Card{Title: "Server Responding", IsOK: true, IconText: "HTTP"},
		Card{Title: "Office2", Metric: "334ms", IsOK: true, ShowHeartPulseIcon: true},
		Card{Title: "Office Synology", Metric: "66ms", IsOK: true, ShowHeartPulseIcon: true},
		Card{Title: "Office Synology Responding", IsOK: true, IconText: "HTTP"},
		Card{Title: "Office3", IsError: true, ShowHeartPulseIcon: true},
		Card{Title: "Home Synology", Metric: "41ms", IsOK: true, ShowHeartPulseIcon: true},
		Card{Title: "Home Synology Responding", IsOK: true, IconText: "HTTP"}}
}
