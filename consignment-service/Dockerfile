FROM golang:1.11 as builder

WORKDIR /go/src/go-micro/consignment-service

COPY . .

RUN ls

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep init -v

RUN dep ensure -v

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .


FROM debian:latest

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/go-micro/consignment-service/consignment-service .

CMD ["./consignment-service"]

