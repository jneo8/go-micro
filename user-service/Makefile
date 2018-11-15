GVM_PROJ_PATH=/home/james/.gvm/pkgsets/go1.11/go-micro/src/go-micro


gen:
	protoc -I. --proto_path=$(GVM_PROJ_PATH)/user-service:. --micro_out=. --go_out=. proto/user/user.proto

build:
	# protoc -I. --go_out=plugins=micro:$(GVM_PROJ_PATH)/consignment-service \
	# 	      proto/consignment/consignment.proto

	protoc -I. --proto_path=$(GVM_PROJ_PATH)/user-service:. --micro_out=. --go_out=. proto/user/user.proto

	# GOOS=linux GOARCH=amd64 go build
	docker build --no-cache -t user-service .

run: 
	docker run --rm -p 50053:50051 \
	    --name user-service \
	    -e MICRO_SERVER_ADDRESS=:50051 \
	    -e MICRO_REGISTRY=mdns user-service

