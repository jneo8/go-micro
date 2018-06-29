gvm_pkgset_path=/home/ubuntu/.gvm/pkgsets/go1.9.7/micro

build:
	protoc -I. --go_out=plugins=grpc:$(gvm_pkgset_path)/src/go-micro/consignment-service proto/consignment/consignment.proto
