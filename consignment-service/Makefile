GVM_PROJ_PATH=/home/ubuntu/.gvm/pkgsets/go1.11/micro/src/go-micro


build:
	protoc -I. --go_out=plugins=grpc:$(GVM_PROJ_PATH)/consignment-service \
		      proto/consignment/consignment.proto

	GOOS=linux GOARCH=amd64 go build
	docker build -t consignment-service .

run: 
	docker run -p 50051:50051 consignment-service
