.PHONY: build dev clean

build: statuspage

dev: statuspage-dev

statuspage: cmd/statuspage/main.go internal/*/* web/*
	go build ./cmd/statuspage

statuspage-dev: cmd/statuspage/main.go internal/*/*
	go build -o statuspage-dev -tags dev ./cmd/statuspage

clean:
	go clean
	rm -f statuspage statuspage-dev
