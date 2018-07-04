build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
	docker build -t consignment-cli .

run:
	docker run -e MICRO_REGISTRY=mdns consignment-cli
