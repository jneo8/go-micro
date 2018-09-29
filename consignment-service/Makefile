GVM_PROJ_PATH=/home/ubuntu/.gvm/pkgsets/go1.11/micro/src/go-micro


build:
	# protoc -I. --go_out=plugins=micro:$(GVM_PROJ_PATH)/consignment-service \
	# 	      proto/consignment/consignment.proto

	protoc -I. --proto_path=$(GVM_PROJ_PATH)/consignment-service:. --micro_out=. --go_out=. proto/consignment/consignment.proto

	# GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run: 
	docker run -p 50051:50051 \
	    -e MICRO_SERVER_ADDRESS=:50051 \
	    -e MICRO_REGISTRY=mdns consignment-service

