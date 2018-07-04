GOCMD=go
gvm_pkgset_path=/home/ubuntu/.gvm/pkgsets/go1.9.7/micro

add_deps:
	$(GOCMD) get github.com/kardianos/govendor
	govendor sync -v
	govendor list

pull_deps:
	govendor sync -v
	govendor list

build:
	# protoc -I. --go_out=plugins=grpc:$(gvm_pkgset_path)/src/go-micro/consignment-service proto/consignment/consignment.proto
	protoc -I. --go_out=plugins=micro:$(gvm_pkgset_path)/src/go-micro/consignment-service \
		proto/consignment/consignment.proto

	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build
	docker build -t consignment-service .

run:
	docker run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns consignment-service
