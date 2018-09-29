GVM_PROJ_PATH=/home/ubuntu/.gvm/pkgsets/go1.11/micro/src/go-micro


build:
	protoc -I. --proto_path=$(GVM_PROJ_PATH)/vessel-service:. --micro_out=. --go_out=. proto/vessel/vessel.proto

	GOOS=linux GOARCH=amd64 go build
	docker build -t vessel-service .

run: 
	docker run -p 50052:50051 \
	    -e MICRO_SERVER_ADDRESS=:50051 \
	    -e MICRO_REGISTRY=mdns vessel-service
