GVM_PROJ_PATH=/home/james/.gvm/pkgsets/go1.11/go-micro/src/go-micro


gen:
	protoc -I. --proto_path=$(GVM_PROJ_PATH)/vessel-service:. --micro_out=. --go_out=. proto/vessel/vessel.proto

build:
	protoc -I. --proto_path=$(GVM_PROJ_PATH)/vessel-service:. --micro_out=. --go_out=. proto/vessel/vessel.proto
	# GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .

run: 
	docker run --rm -p 50052:50051 \
	    --name vessel-service \
	    -e MICRO_SERVER_ADDRESS=:50051 \
	    -e MICRO_REGISTRY=mdns vessel-service

