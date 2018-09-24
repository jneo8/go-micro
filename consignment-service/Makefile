GVM_PROJ_PATH=/home/ubuntu/.gvm/pkgsets/go1.11/micro/src/go-micro


build:
	protoc -I. --go_out=plugins=grpc:$(GVM_PROJ_PATH)/consignment-service \
		      proto/consignment/consignment.proto
